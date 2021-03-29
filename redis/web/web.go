package web

import (
	"github.com/gin-gonic/gin"
	"go-study/redis/server"
)

type Server struct {
	Server *server.Server
}

func NewServer(s *server.Server) *Server {
	return &Server{Server: s}
}

func (s *Server) Run() {
	r := gin.Default()
	s.route(r)
	r.Run(":1234")
}

func (s *Server) route(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	r.GET("/redis", func(c *gin.Context) {
		//clients := make([]string, 0)
		//for _, v := range s.Server.Clients.All() {
		//	clients = append(clients, v.Conn.RemoteAddr().String())
		//}
		//c.JSON(200, gin.H{
		//	"hello":        "world",
		//	"client_count": len(s.Server.Clients.All()),
		//	"clients":      clients,
		//	"storage":      s.Server.Store.All(),
		//})
	})
}
