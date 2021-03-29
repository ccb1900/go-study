package main

import (
	"fmt"
	"go-study/util"
	"time"
)

func main() {
	c := boring("hello")

	// 5 s 后有返回，那个协程最多运行5s
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("you talk too much")
			return
		}
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		util.ILoop(func(i int) {
			c <- fmt.Sprintf("%d %s", i, msg)
			util.Sleep()
		})
	}()

	return c
}
