package server

import (
	"bufio"
	"go-study/redis/exception"
)

type Reply struct {
}

func NewReply() *Reply {
	return &Reply{}
}

func (r *Reply) Write(w *bufio.Writer, st string) {
	if _, err := w.Write([]byte(st)); err != nil {
		exception.Report(err, "write  error..")
	}

	if err := w.Flush(); err != nil {
		exception.Report(err, "flush failed")
	}
}
