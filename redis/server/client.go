package server

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"ppp/redis/exception"
	"ppp/redis/packet"
	"strconv"
	"strings"
)

type Client struct {
	Conn      net.Conn
	Id        int
	DBNum     int
	BufReader *bufio.Reader
	BufWriter *bufio.Writer
}

func NewClient(c net.Conn, id int) *Client {
	client := new(Client)
	client.Conn = c
	client.Id = id
	client.DBNum = 0
	return client
}

func (c *Client) ParsePacket() ([]string, string, error) {
	line, isPrefix, err := c.BufReader.ReadLine()
	exception.Debug(line, isPrefix, err)
	rawCommand := ""
	pc := 0
	if strings.HasPrefix(string(line), "*") {
		pc, err = strconv.Atoi(string(line[1:]))
		rawCommand += string(line) + packet.EL
		if err != nil {
			return nil, "", errors.New(fmt.Sprintf("unknown command `%s`, with args beginning with: ", string(line)))
		}
	}
	// 记录命令
	command := make([]string, 0)
	for i := 0; i < pc; i++ {
		line, isPrefix, err = c.BufReader.ReadLine()
		if err != nil {
			return nil, "", err
		}
		rawCommand += string(line) + packet.EL
		if strings.HasPrefix(string(line), "$") {
			line, isPrefix, err = c.BufReader.ReadLine()
			if err != nil {
				return nil, "", err
			}
			rawCommand += string(line) + packet.EL
			command = append(command, string(line))
		}
	}
	return command, rawCommand, nil
}
