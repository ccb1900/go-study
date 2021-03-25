package effective_go

func incr() {

}

func main() {
	cnt := 0
	c := make(chan int)
	c <- cnt
	for i := 0; i < 10; i++ {
		go func() {
			cnt = <-c
			cnt++
			c <- cnt
		}()
	}

}
