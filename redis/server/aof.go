package server

import "os"

type Aof struct {
	File   *os.File
	AofBuf chan string
}

func NewAof() *Aof {
	f, _ := os.OpenFile("redis.aof", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	return &Aof{
		File:   f,
		AofBuf: make(chan string, 1024),
	}
}

func (a *Aof) Save(s string) {
	a.File.Write([]byte(s))
	a.File.Sync()
}
