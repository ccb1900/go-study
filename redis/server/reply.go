package server

import "bufio"

type Reply struct {
	Key []string
	W   *bufio.Writer
}

func NewReply(k []string, w *bufio.Writer) *Reply {
	return &Reply{
		Key: k,
		W:   w,
	}
}
