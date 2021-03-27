package server

import "bufio"

type Command struct {
	Commands []string
	DbNum    int
	Writer   *bufio.Writer
}

func NewCommand(commands []string, dbNum int, writer *bufio.Writer) *Command {
	return &Command{
		Commands: commands,
		DbNum:    dbNum,
		Writer:   writer,
	}
}

func (c *Command) validate() bool {
	return false
}
