package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	path := "/Users/ppandey/Desktop/concurrencyingolang/TextFiles"
	data, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("error while getting dir", err)
		return
	}

	job := make(chan string, len(data))
	result := make(chan map[string]int, len(data))

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
		go Worker4(i, job, result)
	}

	// sending value in job // sending job
	for _, val := range s {
		job <- val
	}
	close(job)
	for m := range result {
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

}

func Worker4(workerId int, job chan string, result chan map[string]int) {
	for file := range job {
		workerMap := map[string]int{}
		str := tokenize1(file)
		for _, val := range str {
			workerMap[val]++
		}
		result <- workerMap
	}
}

func tokenize1(text string) []string {
	var words []string
	var sb strings.Builder

	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sb.WriteRune(unicode.ToLower(r))
		} else if sb.Len() > 0 {
			words = append(words, sb.String())
			sb.Reset()
		}
	}

	if sb.Len() > 0 {
		words = append(words, sb.String())
	}

	return words
}
