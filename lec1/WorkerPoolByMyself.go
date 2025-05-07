package main

import "fmt"

func main() {
	job := make(chan int, 100)
	result := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go Worker(i, job, result)
	}

	for i := 1; i <= 100; i++ {
		job <- i
	}
	close(job)

	for val := range result {
		fmt.Println("val:", val)
	}
	close(result)
}

func Worker(id int, job chan int, result chan int) {
	fmt.Println("Worker:", id)
	for val := range job {
		result <- val * val
	}
}
