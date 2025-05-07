package main

import (
	"fmt"
	"time"
)

func main() {
	ans := 0
	t := time.Now()
	for i := 1; i <= 10000; i++ {
		ans += i * i
	}
	fmt.Println(ans)
	fmt.Println("Took this much time,", time.Since(t))
}

//333383335000
//Took this much time, 84.917µs
//
//took this much time  1.880875ms
