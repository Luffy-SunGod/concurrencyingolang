package main

import (
	"context"
	"fmt"
	"time"
)

func workerw(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped!")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go workerw(ctx)

	time.Sleep(2 * time.Second) // Thoda kaam karne do
	fmt.Println("Main: Cancelling context now!")
	cancel() // Worker ko bol diya, bas ab band ho jao

	time.Sleep(1 * time.Second)
	fmt.Println("Main: Exiting...")
}
