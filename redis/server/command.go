package server

import "bufio"

type Command struct {
	Commands   []string
	Writer     *bufio.Writer
	Client     *Client
	RawCommand string
}

func NewCommand(cc *Client, commands []string, rawCommand string, writer *bufio.Writer) *Command {
	return &Command{
		Commands:   commands,
		Writer:     writer,
		Client:     cc,
		RawCommand: rawCommand,
	}
}

func (c *Command) validate() bool {
	return true
}
