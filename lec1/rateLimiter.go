package main

import (
	"fmt"
	"time"
)

func main() {
	// 5 requests per second allowed
	rateLimit := 5
	limiter := time.NewTicker(time.Second * 5 / time.Duration(rateLimit)) // Create a ticker for rate limiting

	// Channels for task queue and response
	taskQueue := make(chan int, 10) // Task queue with some tasks to process
	done := make(chan struct{})     // Done channel to indicate when all tasks are processed

	// Simulate adding tasks
	go func() {
		fmt.Println("I amadding value in task")
		for i := 1; i <= 10; i++ {
			taskQueue <- i
			fmt.Println("Added task:", i)
		}
		close(taskQueue)
		fmt.Println("Completed value addition")
	}()

	// Process tasks with rate limiting
	go func() {
		fmt.Println("Now getting values")
		for task := range taskQueue {
			<-limiter.C // Wait for the limiter to allow a task to process
			fmt.Printf("Processing task %d at %v\n", task, time.Now())
		}
		done <- struct{}{}
		fmt.Println(" getting values completed")
	}()
	fmt.Println("Till now main thread is not blockedQ!!")
	<-done // Wait until all tasks are processed
	fmt.Println("All tasks processed.")
}
