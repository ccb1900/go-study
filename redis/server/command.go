package server

import (
	"strings"
)

type Command struct {
	Commands   []string
	Client     *Client
	RawCommand string
}

func NewCommand(cc *Client, commands []string, rawCommand string) *Command {
	return &Command{
		Commands:   commands,
		Client:     cc,
		RawCommand: rawCommand,
	}
}

func (c *Command) validate() bool {
	// 是否为空
	l := len(c.Commands)
	if l == 0 {
		return false
	}

	// 是否为单目命令
	whiteList := []string{
		"ping",
	}

	cd := strings.ToLower(c.Commands[0])
	for _, wl := range whiteList {
		if cd == wl {
			return cd == wl
		}
	}

	switch cd {
	case "set":
		if l < 3 {
			return false
		}
	}
	return true
}
