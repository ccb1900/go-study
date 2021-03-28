package effective_go

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func myIf(x int) int {
	if x > 0 {
		return x + 1
	} else {
		return x - 1
	}
}

func ctrlError(err error) {
	if err := err; err != nil {
		fmt.Println(err)
	}
}

func loop(n int) {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func rangeChannel(ch chan string) {
	for c := range ch {
		fmt.Println(c)
	}
}
func mergeCh(c ...chan string) chan string {
	cc := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(c))
	for _, v := range c {
		go func(v chan string) {
			for k := range v {
				cc <- k
			}
			wg.Done()
		}(v)
	}
	go func() {
		wg.Wait()
		close(cc)
	}()
	return cc
}
func rangeMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func rangeSlice(s []string) {
	for k, v := range s {
		fmt.Println(k, v)
	}
}

func wCh() chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			ch <- strconv.Itoa(i) + "hello"
		}
		close(ch)
	}()
	return ch
}
