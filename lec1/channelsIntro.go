package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		v := <-c //recieving value
		fmt.Println(v)
	}() //sending value to channels
	c <- 5

}
