package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	n := 10

	for i := 0; i < n; i++ {
		go func() {
			time.Sleep(time.Second)
		}()
	}
	fmt.Println("架构", runtime.GOARCH)
	fmt.Println("架构", runtime.GOROOT())
	fmt.Println("架构", runtime.GOOS)
	fmt.Println("协程数量", runtime.NumGoroutine())
	fmt.Println("架构", runtime.Compiler)
	fmt.Println("架构", runtime.Version())
	fmt.Println("架构", runtime.NumCPU())
}
