package main

import (
	"fmt"
	"ppp/redis/server"
)

func main() {
	s := server.NewServer(":9009")
	s.Run()
	fmt.Println(s)
}
