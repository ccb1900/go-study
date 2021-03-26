package server

import "net"

type Client struct {
	Conn  net.Conn
	Id    int
	DBNum int
}

func NewClient(c net.Conn, id int) *Client {
	client := new(Client)
	client.Conn = c
	client.Id = id
	client.DBNum = 0
	return client
}
