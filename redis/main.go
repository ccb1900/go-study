package main

import (
	"fmt"
	"go-study/redis/server"
	"go-study/redis/web"
	"log"
	"net/http"
	"sync"
)
import _ "net/http/pprof"

func main() {
	s := server.NewServer(":9009")
	w := web.NewServer(s)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		s.Run()
		wg.Done()
	}()

	go func() {
		w.Run()
		wg.Done()
	}()

	go func() {
		log.Println(http.ListenAndServe(":1267", nil))
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("start...")
}
