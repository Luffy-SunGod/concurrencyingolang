package main

import (
	"fmt"
	"os"
)

func main() {
	path := "/Users/ppandey/Desktop/concurrencyingolang/TextFiles"
	data, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("error while getting dir", err)
		return
	}

	job := make(chan string, len(data))
	result := make(chan int, len(data))

	s := []string{}

	for _, val := range data {
		if val.IsDir() {
			fmt.Println("Skipping its a directory!!")
			continue
		}
		filePath := path + "/" + val.Name()
		d2, err := os.ReadFile(filePath)
		if err == nil {
			s = append(s, string(d2))
		}

	}

	for i := 1; i < 3; i++ {
		go Worker3(i, job, result)
	}

	// sending value in job // sending job
	for _, val := range s {
		job <- val
		fmt.Println("file size is", <-result)
	}
}

func Worker3(workerId int, job chan string, result chan int) {
	for file := range job {
		fmt.Println("I am go routine no :", workerId)
		result <- len(file)
	}
}
