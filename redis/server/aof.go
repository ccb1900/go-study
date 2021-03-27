package server

import "os"

type Aof struct {
	File   *os.File
	AofBuf chan string
}
