package main

import (
	"fmt"
	"go-study/util"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

}

func workerEfficient(id int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup
	for job := range jobs {
		wg.Add(1)

		go func(j int, id int) {
			fmt.Println("worker", id, "started job", j)
			time.Sleep(time.Second)
			fmt.Println("worker", id, "fnished job", j)
			results <- j * 2
			wg.Done()
		}(job, id)
	}
	wg.Wait()
}

func main() {
	numJobs := 8

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	util.Loop(func(i int, n int) {
		go workerEfficient(i+1, jobs, results)
	}, 3)

	util.Loop(func(i int, n int) {
		jobs <- i + 1
	}, numJobs)

	close(jobs)

	util.Loop(func(i int, n int) {
		<-results
	}, numJobs)

	close(results)
}
