package main

import (
	"fmt"
	"ppp/redis/server"
	"ppp/redis/web"
	"sync"
)

func main() {
	s := server.NewServer(":9009")
	w := web.NewServer(s)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		s.Run()
		wg.Done()
	}()

	go func() {
		w.Run()
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("start...")
}
