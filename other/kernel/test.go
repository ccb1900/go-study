package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	cc := make(chan int)
	cc2 := make(chan int)
	go func() {
		a := rand.Intn(10)
		b := rand.Intn(10)

		c := a + b
		fmt.Println(a, b, c)
		cc <- 1
		close(cc)
	}()

	go func() {
		a := rand.Intn(10)
		b := rand.Intn(10)
		c := a + b
		fmt.Println(a, b, c)

		cc2 <- 1
		close(cc2)
	}()

	// 聚合结果
	c := make(chan int)
	w := sync.WaitGroup{}
	w.Add(2)
	go func() {
		for ccc := range cc2 {
			c <- ccc
		}
		defer w.Done()
	}()
	go func() {
		for ccc := range cc {
			c <- ccc
		}
		defer w.Done()
	}()

	go func() {
		for mm := range c {
			fmt.Println("sss", mm)
		}
	}()
	w.Wait()
	close(c)
	fmt.Println("complete")
}
