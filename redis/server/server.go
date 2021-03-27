package server

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"net"
	"ppp/redis/exception"
	"ppp/redis/packet"
	"ppp/redis/snowflake"
	"strconv"
	"strings"
)

type Server struct {
	DBList      []*DB
	Address     string
	Listener    net.Listener
	Clients     map[int]*Client
	RemoveList  chan int
	CommandList chan *Command
	WriteList   chan *Reply
	ClientList  chan *Client
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
		case ol := <-s.CommandList:
			if !ol.validate() {
				go s.Resp(ol.Writer, packet.ErrLine("ERR syntax error"))
			} else {
				switch strings.ToLower(ol.Commands[0]) {
				case "set":
					s.SetKey(ol.DbNum, ol.Commands[1], NewObject(ol.DbNum, ol.Commands[1], ol.Commands[2]))
					go s.Resp(ol.Writer, packet.OkLine("OK"))
				case "get":
					if o, err := s.GetObject(ol.DbNum, ol.Commands[1]); err != nil {
						go s.Resp(ol.Writer, packet.OkLine("(nil)"))
					} else {
						go s.Resp(ol.Writer, packet.GetString(o.Value))
					}

				default:
					go s.Resp(ol.Writer, packet.OkLine("OK"))
				}
				exception.Debug("storage client")
			}

		}
	}
}

// 接收客户端连接
func (s *Server) handle(c *Client) {
	w := bufio.NewWriter(c.Conn)
	bufReader := bufio.NewReader(c.Conn)
	for {
		exception.Debug("read start")
		// 等待连接继续发命令
		// 开始读命令
		command := s.parsePacket(bufReader, w)

		// 读完一次发送的包
		if len(command) == 0 {
			exception.Debug("客户端断开" + c.Conn.RemoteAddr().String())
			go s.closeClient(c)
			break
		}
		go s.handleCommand(w, command, c)
	}
}
func (s *Server) closeClient(c *Client) {
	s.RemoveList <- c.Id
}
func (s *Server) handleCommand(w *bufio.Writer, commands []string, c *Client) {
	exception.Debug("解析出来的命令", commands)
	s.CommandList <- NewCommand(commands, c.DBNum, w)
}

// 根据协议解析命令
func (s *Server) parsePacket(bufReader *bufio.Reader, w *bufio.Writer) []string {
	line, isPrefix, err := bufReader.ReadLine()
	exception.Debug(line, isPrefix, err)
	pc := 0
	if strings.HasPrefix(string(line), "*") {
		pc, err = strconv.Atoi(string(line[1:]))
		if err != nil {
			s.Failed(w, err, "转换出错")
		}
	}
	// 记录命令
	command := make([]string, 0)
	for i := 0; i < pc; i++ {
		line, isPrefix, err = bufReader.ReadLine()
		if err != nil {
			s.Failed(w, err, "读取参数数量出错")
		}
		if strings.HasPrefix(string(line), "$") {
			line, isPrefix, err = bufReader.ReadLine()
			if err != nil {
				s.Failed(w, err, "读取参数出错")
			}
			command = append(command, string(line))
		}
	}
	return command
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
	s.Resp(w, packet.ErrLine(st))
	exception.Report(err, st)
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
	for i := 0; i < 16; i++ {
		s.DBList[i] = NewDB(i)
	}
	return s
}
