package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := fanInSimple(boring("hello"), boring("world"))

	for k := 0; k < 50; k++ {
		fmt.Println(<-c)
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for k := 0; ; k++ {
			c <- fmt.Sprintf("%s %d", msg, k)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			v1 := <-c1
			c <- v1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}

func fanInSimple(cs ...<-chan string) <-chan string {
	c := make(chan string)

	for _, ci := range cs {
		go func(cv <-chan string) {
			for {
				c <- <-cv
			}
		}(ci)
	}

	return c
}
