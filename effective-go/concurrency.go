package effective_go

import (
	"fmt"
	"sort"
)

func incr() {

}

func main() {
	l := []int{11, 22, 33, 44}

	c := make(chan bool)
	cl := make(chan int, 4)
	go func() {
		sort.Ints(l)
		for _, v := range l {
			cl <- v
		}
		close(cl)
		c <- true
	}()
	<-c
	for cc := range cl {
		fmt.Println("this is ", cc)
	}
	fmt.Println(l)

}
