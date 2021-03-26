package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

type Server struct {
	Address  string
	Listener net.Listener
}

func (s *Server) Run() {
	for {
		log.Println("server is waiting...")
		c, err := s.Listener.Accept()
		if err != nil {
			log.Fatalf("client accept error..%v", err)
		}
		log.Println("client arrive at your server..")
		go s.handle(c)
	}
}
func (s *Server) handle(c net.Conn) {
	fmt.Println(c)
	n, err := c.Write([]byte("hello,world"))
	if err != nil {
		log.Fatalf("handle client  error..%v", err)
	}

	fmt.Println("n->", n)
}
func NewServer(address string) *Server {
	s := new(Server)
	s.Address = address
	l, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Fatalf("start %v", err)
	}
	s.Listener = l

	return s
}
