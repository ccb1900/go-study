package main

import (
	"fmt"
	"go-study/util"
	"sync"
)

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		util.Loop(func(i int, n int) {
			out <- fmt.Sprintf("配件 %d", i+1)
		}, n)
	}()

	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for i := range in {
			out <- fmt.Sprintf("组装(%s)", i)
			util.SleepN(1)
		}
	}()

	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for i := range in {
			out <- fmt.Sprintf("打包(%s)", i)
		}
	}()

	return out
}

func merge(items ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(items))
	for _, item := range items {
		go func(item <-chan string) {
			defer wg.Done()
			for x := range item {
				out <- x
			}
		}(item)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	items := buy(10)
	phones := merge(build(items), build(items), build(items), build(items), build(items))
	packs := pack(phones)

	for p := range packs {
		fmt.Println(p)
	}
}
