package effective_go

import (
	"fmt"
	"sort"
)

func incr() {

}

func main() {
	l := []int{1, 2, 3, 4}

	c := make(chan bool)
	go func() {
		sort.Ints(l)
		c <- true
	}()
	<-c
	fmt.Println(l)

}
