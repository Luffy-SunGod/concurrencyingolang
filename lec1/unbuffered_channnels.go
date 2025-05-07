package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 1
	c <- 2
	//c <- 3 // buffer is full hence its blocking

	x := <-c
	fmt.Println(x)
	x = <-c
	fmt.Println(x)

	//buffer is empty hence its blocking
	x = <-c
	fmt.Println(x)

}
