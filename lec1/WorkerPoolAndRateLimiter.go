package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	job := make(chan int, 100)
	response := make(chan int, 100)
	wg := &sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(i)
		go SendRequestWorker(i, job, response, wg)
	}
	for i := 1; i <= 100; i++ {
		job <- i
	}
	close(job)
	go func() {
		wg.Wait()
		close(response)
	}()

	for val := range response {
		fmt.Println(val)
	}

}

func SendRequestWorker(workerId int, job chan int, response chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	rateLimit := 3
	for val := range job {
		limiter := time.NewTicker(time.Second / time.Duration(rateLimit))
		<-limiter.C
		fmt.Println("sending request!!!", val)
		response <- val
	}

}
