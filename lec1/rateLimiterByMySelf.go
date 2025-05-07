package main

import (
	"fmt"
	"time"
)

func main() {
	req := make(chan int, 10)
	notifyChan := make(chan struct{})

	rateLimit := 5
	limiter := time.NewTicker(time.Second / time.Duration(rateLimit))

	go func() {
		for i := 1; i <= 20; i++ {
			req <- i
			fmt.Println("Sended req :", i)
		}
		close(req)
	}()

	go func(req chan int) {
		for r := range req {
			fmt.Println(<-limiter.C)
			fmt.Println("executing request ", r)
		}

		notifyChan <- struct{}{}
	}(req)
	<-notifyChan
	fmt.Println("Programe ended!!")

}
