package main

import (
	"context"
	"fmt"
)

type Request struct {
	Name string
}

func test(ctx context.Context) error {
	for {
		v := demo(ctx)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case vv := <-v:
			fmt.Println(vv)
		}
	}
	//ctx.Value(Request{Name: "hello"})
	//x, _ := context.WithCancel(ctx)
	//y, _ := context.WithTimeout(x, time.Second)
	//z := context.WithValue(y, Request{Name: "hello"}, Request{Name: "world"})
	//w, _ := context.WithDeadline(z, time.Now())
	//fmt.Println(x, y, z, w)
}
func demo(ctx context.Context) chan int {
	c := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case c <- n:
				n++
			}
		}
	}()

	return c
}

func main() {
	ctx := context.Background()
	test(ctx)
	demo(ctx)
	fmt.Println("he")
}
