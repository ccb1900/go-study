package server

import (
	log "github.com/sirupsen/logrus"
	"net"
	"ppp/redis/exception"
)

type Net struct {
	Address  string
	Listener net.Listener
}

func NewNet(addr string) *Net {
	l, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("start %v", err)
	}
	return &Net{
		Address:  addr,
		Listener: l,
	}
}

func (n *Net) Accept() net.Conn {
	exception.Debug("server is waiting...")
	c, err := n.Listener.Accept()
	if err != nil {
		log.Fatalf("client accept error..%v", err)
	}
	exception.Debug("client arrive at your server.." + c.RemoteAddr().String())

	return c
}
