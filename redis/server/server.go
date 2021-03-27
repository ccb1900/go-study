package server

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"ppp/redis/exception"
	"ppp/redis/packet"
	"ppp/redis/snowflake"
	"strconv"
	"strings"
)

type Server struct {
	DBList      []*DB
	DbNum       int
	Address     string
	Listener    net.Listener
	Clients     map[int]*Client
	RemoveList  chan int
	CommandList chan *Command
	WriteList   chan *Reply
	ClientList  chan *Client
	RDBFile     *os.File
	Aof         *Aof
	Rdb         *Rdb
}

// 创建redis server
func NewServer(address string) *Server {
	s := new(Server)
	s.Address = address
	l, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Fatalf("start %v", err)
	}
	s.Listener = l
	s.Clients = make(map[int]*Client, 128)
	s.RemoveList = make(chan int, 128)
	s.CommandList = make(chan *Command, 1024)
	s.WriteList = make(chan *Reply, 1024)
	s.ClientList = make(chan *Client, 128)
	s.DBList = make([]*DB, 16)
	s.Aof = NewAof()
	s.DbNum = 16

	for i := 0; i < 16; i++ {
		s.DBList[i] = NewDB(i)
	}
	return s
}
func (s *Server) SetKey(dbNum int, k string, v *Object) {
	s.DBList[dbNum].Store[k] = v
}

func (s *Server) GetObject(dbNum int, k string) (*Object, error) {
	if v, exit := s.DBList[dbNum].Store[k]; exit {
		return v, nil
	}

	return nil, errors.New("key not exists")

}

// 运行server
func (s *Server) Run() {
	go func() {
		for {
			exception.Debug("server is waiting...")
			c, err := s.Listener.Accept()
			if err != nil {
				log.Fatalf("client accept error..%v", err)
			}
			exception.Debug("client arrive at your server.." + c.RemoteAddr().String())
			clientId := snowflake.GetId()
			client := NewClient(c, clientId)
			go func(client *Client) {
				s.ClientList <- client
			}(client)
			go s.handle(client)
		}
	}()
	for {
		select {
		case cl := <-s.ClientList:
			s.Clients[cl.Id] = cl
			exception.Debug("range client")
		case rl := <-s.RemoveList:
			delete(s.Clients, rl)
			exception.Debug("delete client")
		case aof := <-s.Aof.AofBuf:
			go s.Aof.Save(aof)
		case ol := <-s.CommandList:
			if !ol.validate() {
				go s.Resp(ol.Client.BufWriter, packet.ErrLine("ERR syntax error"))
			} else {
				switch strings.ToLower(ol.Commands[0]) {
				case "set":
					s.SetKey(ol.Client.DBNum, ol.Commands[1], NewObject(ol.Client.DBNum, ol.Commands[1], ol.Commands[2]))
					go s.Resp(ol.Client.BufWriter, packet.OkLine("OK"))
					go func(ccc string) {
						s.Aof.AofBuf <- ccc
					}(ol.RawCommand)
				case "get":
					if o, err := s.GetObject(ol.Client.DBNum, ol.Commands[1]); err != nil {
						go s.Resp(ol.Client.BufWriter, packet.OkLine("(nil)"))
					} else {
						go s.Resp(ol.Client.BufWriter, packet.GetString(o.Value))
					}
				case "select":
					dbNum, _ := strconv.Atoi(ol.Commands[1])
					ol.Client.DBNum = dbNum % s.DbNum
					go s.Resp(ol.Client.BufWriter, packet.OkLine("OK"))
					go func(ccc string) {
						s.Aof.AofBuf <- ccc
					}(ol.RawCommand)
				case "ping":
					if len(ol.Commands) == 2 {
						go s.Resp(ol.Client.BufWriter, packet.GetString(ol.Commands[1]))
					} else {
						go s.Resp(ol.Client.BufWriter, packet.OkLine("PONG"))
					}
				case "client":
					if ol.Commands[1] == "list" {
						go s.Resp(ol.Client.BufWriter, packet.OkLine(fmt.Sprintf("id=%d addr=%s db=%d", ol.Client.Id, ol.Client.Conn.RemoteAddr().String(), ol.Client.DBNum)))
					} else {
						go s.Resp(ol.Client.BufWriter, packet.ErrLine(fmt.Sprintf("ERR Unknown subcommand or wrong number of arguments for '%s'. Try CLIENT HELP", ol.Commands[1])))
					}
				default:
					go s.Resp(ol.Client.BufWriter, packet.OkLine("OK"))
				}

				exception.Debug("storage client")
			}

		}
	}
}

// 接收客户端连接
func (s *Server) handle(c *Client) {
	c.BufReader = bufio.NewReader(c.Conn)
	c.BufWriter = bufio.NewWriter(c.Conn)
	for {
		exception.Debug("read start")
		// 等待连接继续发命令
		// 开始读命令
		cd, err := c.ParsePacket()
		if err == io.EOF {
			s.RemoveList <- c.Id
			break
		}
		if err != nil {
			s.Failed(c.BufWriter, err, err.Error())
		} else {
			s.CommandList <- cd
		}
		// 读完一次发送的包
		//if len(command) == 0 {
		//	exception.Debug("客户端断开" + c.Conn.RemoteAddr().String())
		//	go s.closeClient(c)
		//	break
		//}

	}
}
func (s *Server) closeClient(c *Client) {
	s.RemoveList <- c.Id
	c.Conn.Close()
}

// 成功响应
func (s *Server) Resp(w *bufio.Writer, st string) {
	if _, err := w.Write([]byte(st)); err != nil {
		exception.Report(err, "write  error..")
	}

	if err := w.Flush(); err != nil {
		exception.Report(err, "flush failed")
	}
}

// 返回错误信息
func (s *Server) Failed(w *bufio.Writer, err error, st string) {
	exception.Report(err, st)
	s.Resp(w, packet.ErrLine(st))
}
