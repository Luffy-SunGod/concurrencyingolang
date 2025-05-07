package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker function — ye alag goroutine se data bhejega
func worker(id int, out chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		out <- id*100 + i
	}
}

// Fan-In function — multiple input channels ko merge karega
func fanIn(cs ...<-chan int) <-chan int {
	out := make(chan int)

	for _, c := range cs {
		go func(ch <-chan int) {
			for v := range ch {
				out <- v
			}
		}(c)
	}

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Teen alag workers, teen channels
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go worker(1, c1)
	go worker(2, c2)
	go worker(3, c3)

	// Fan-In — combine kar diya
	out := fanIn(c1, c2, c3)

	// Read combined output
	for i := 0; i < 15; i++ {
		fmt.Println("Received:", <-out)
	}

	fmt.Println("All done")
}
