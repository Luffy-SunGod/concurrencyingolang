package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	job := make(chan int, 100)
	result := make(chan int)
	t := time.Now()

	// making 3 workers
	wg := &sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, job, result, wg)
	}

	for i := 1; i <= 10000; i++ {
		job <- i
	}
	close(job)

	go func() {
		wg.Wait()
		close(result)
	}()
	ans := 0

	for val := range result {
		ans += val
	}
	fmt.Println(ans)
	fmt.Println("took this much time ", time.Since(t))

}

func worker(id int, job chan int, result chan int, wg *sync.WaitGroup) {
	ans := 0
	defer wg.Done()
	for val := range job {
		ans += val * val
	}
	result <- ans
}
