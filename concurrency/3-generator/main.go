package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := boring("hello")
	b := boring("world")

	for i := 0; i < 12; i++ {
		fmt.Println(<-a)
		fmt.Println(<-b)
	}

	fmt.Println("i am leving")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		close(c)
	}()

	return c
}
