package server

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"ppp/redis/exception"
	"ppp/redis/packet"
	"ppp/redis/snowflake"
	"strconv"
	"strings"
)

type Server struct {
	DBList       *DBCollection
	Clients      map[int]*Client
	RDBFile      *os.File
	Aof          *Aof
	Rdb          *Rdb
	ErrorMessage *Message
	Sync         *Sync
	Reply        *Reply
	Net          *Net
}

// 创建redis server
func NewServer(address string) *Server {
	s := new(Server)
	s.Net = NewNet(address)
	s.Clients = make(map[int]*Client, 128)
	s.Sync = NewSync()
	s.DBList = NewDbList(16)
	s.Aof = NewAof()
	s.Reply = NewReply()
	s.ErrorMessage = NewMessage()
	return s
}

func (s *Server) acceptHandler() {
	for {
		client := NewClient(s.Net.Accept(), snowflake.GetId())
		s.Sync.ClientList <- client
		go s.handle(client)
	}
}

func (s *Server) Loop() {
	for {
		select {
		case cl := <-s.Sync.ClientList:
			s.Clients[cl.Id] = cl
			exception.Debug("range client")
		case rl := <-s.Sync.RemoveList:
			delete(s.Clients, rl)
			exception.Debug("delete client")
		case aof := <-s.Aof.AofBuf:
			go s.Aof.Save(aof)
		case ol := <-s.Sync.CommandList:
			if !ol.validate() {
				go s.Resp(ol.Client.BufWriter, packet.ErrLine("ERR syntax error"))
			} else {
				switch strings.ToLower(ol.Commands[0]) {
				case "set":
					s.DBList.Set(ol.Client.DBNum, ol.Commands[1], NewObject(ol.Client.DBNum, ol.Commands[1], ol.Commands[2]))
					go s.Resp(ol.Client.BufWriter, packet.OkLine("OK"))
					go func(ccc string) {
						s.Aof.AofBuf <- ccc
					}(ol.RawCommand)
				case "get":
					if o, err := s.DBList.Get(ol.Client.DBNum, ol.Commands[1]); err != nil {
						go s.Resp(ol.Client.BufWriter, packet.OkLine("(nil)"))
					} else {
						go s.Resp(ol.Client.BufWriter, packet.GetString(o.Value))
					}
				case "select":
					dbNum, _ := strconv.Atoi(ol.Commands[1])
					ol.Client.DBNum = dbNum % s.DBList.Num
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

// 运行server
func (s *Server) Run() {
	go s.acceptHandler()
	s.Loop()
}

// 接收客户端连接
func (s *Server) handle(c *Client) {
	c.BufReader = bufio.NewReader(c.Conn)
	c.BufWriter = bufio.NewWriter(c.Conn)
	for {
		exception.Debug("read start")
		cd, err := c.ParsePacket()
		// 客户端断开连接
		if err == io.EOF {
			s.Sync.RemoveList <- c.Id
			break
		}
		if err != nil {
			s.Failed(c.BufWriter, err, err.Error())
		} else {
			s.Sync.CommandList <- cd
		}
	}
}
func (s *Server) closeClient(c *Client) {
	s.Sync.RemoveList <- c.Id
	c.close()
}

// 成功响应
func (s *Server) Resp(w *bufio.Writer, st string) {
	s.Reply.Write(w, st)
}

// 返回错误信息
func (s *Server) Failed(w *bufio.Writer, err error, st string) {
	exception.Report(err, st)
	s.Reply.Write(w, st)
}
