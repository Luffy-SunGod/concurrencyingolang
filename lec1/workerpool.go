package main

import (
	"fmt"
	"time"
)

func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("👷 Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second) // simulate time-consuming task
		results <- j * j
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// 3 worker goroutines
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channel since all jobs are sent

	// Receive all results
	for a := 1; a <= 5; a++ {
		fmt.Println("✅ Result:", <-results)
	}
}
