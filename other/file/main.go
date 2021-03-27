package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	//_ = ioutil.WriteFile("test.txt", []byte("hello"), 0644)
	f, e := os.OpenFile("text.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	log.Println(f, e)
	n, e := f.Write([]byte("hello"))
	log.Println(n, e)
	f.Sync()
}
