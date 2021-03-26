package server

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"net"
	"ppp/redis/exception"
	"ppp/redis/packet"
	"ppp/redis/snowflake"
	"strconv"
	"strings"
)

type Server struct {
	Address    string
	Listener   net.Listener
	Clients    map[int]*Client
	Store      map[string]*Object
	RemoveList chan int
	ObjectList chan *Object
	WriteList  chan *Reply
}

// 运行server
func (s *Server) Run() {
	clientChan := make(chan *Client, 100)
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
				clientChan <- client
			}(client)
			go s.handle(client)
		}
	}()
	for {
		select {
		case cl := <-clientChan:
			s.Clients[cl.Id] = cl
			exception.Debug("range client")
		case rl := <-s.RemoveList:
			delete(s.Clients, rl)
			exception.Debug("delete client")
		case ol := <-s.ObjectList:
			s.Store[ol.Key] = ol
			exception.Debug("storage client")
		case wl := <-s.WriteList:
			exception.Debug("resp::")
			if len(wl.Key) > 0 {
				switch wl.Key[0] {
				case "get":
					s.Resp(wl.W, packet.OkLine("\""+s.Store[wl.Key[1]].Value+"\""))
				case "command":
					s.Resp(wl.W, packet.OkLine("OK"))
				case "config":
					s.Resp(wl.W, packet.ErrLine("OK"))
				default:
					s.Resp(wl.W, packet.OkLine("OK"))
				}
			} else {
				s.Resp(wl.W, packet.OkLine("OK"))
			}

			exception.Debug("fetch client")
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
			s.closeClient(c)
			break
		}
		go func(command []string) {
			exception.Debug("解析出来的命令", command)
			s.handleCommand(w, command)
		}(command)
	}
}
func (s *Server) closeClient(c *Client) {
	s.RemoveList <- c.Id
}
func (s *Server) handleCommand(w *bufio.Writer, commands []string) {
	// 很多个协程可以写map
	switch strings.ToLower(commands[0]) {
	case "set":
		s.ObjectList <- NewObject(commands[1], commands[2])
		s.WriteList <- NewReply(make([]string, 0), w)
		return
	case "get":
		s.WriteList <- NewReply(commands, w)
		return
	case "command":
		s.WriteList <- NewReply(make([]string, 0), w)
		return
	default:
		s.WriteList <- NewReply(make([]string, 0), w)
		return
	}
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
	s.Store = make(map[string]*Object)
	s.Clients = make(map[int]*Client)
	s.RemoveList = make(chan int)
	s.ObjectList = make(chan *Object)
	s.WriteList = make(chan *Reply)
	return s
}
