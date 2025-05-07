package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"unicode"
)

func main() {
	path := "/Users/ppandey/Desktop/concurrencyingolang/TextFiles"

	fmt.Println("I am line 16!")
	files, err := os.ReadDir(path)
	fmt.Println("I am line 18!")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Prepare input channel and output channel
	jobChan := make(chan string)
	resultChan := make(chan map[string]int)

	// Fan-out: 3 workers
	var wg sync.WaitGroup
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker5(i, jobChan, resultChan, &wg)
	}

	// Feed jobs
	go func() {
		fmt.Println("I m  feeder!!")
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			filePath := filepath.Join(path, f.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("error reading file:", filePath, err)
				continue
			}
			jobChan <- string(data)
		}
		close(jobChan)
	}()

	// Close resultChan once all workers are done
	go func() {
		fmt.Println("I am waiter!!")
		wg.Wait()
		close(resultChan)
	}()
	fmt.Println("I am line 60")
	// Fan-in: Combine results
	finalMap := map[string]int{}
	for partialMap := range resultChan {
		fmt.Println("I am printing maps")
		for word, count := range partialMap {
			finalMap[word] += count
		}
	}

	// Show top 5 words
	printTopWords(finalMap, 5)
}

// Worker function
func worker5(id int, jobs <-chan string, results chan<- map[string]int, wg *sync.WaitGroup) {
	fmt.Println("I am worker!!", id)
	defer wg.Done()
	for text := range jobs {
		counts := make(map[string]int)
		for _, word := range tokenize(text) {
			counts[word]++
		}
		results <- counts
	}
}

// Tokenizer function
func tokenize(text string) []string {
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

// Print top N words
func printTopWords(freq map[string]int, top int) {
	type wordCount struct {
		word  string
		count int
	}
	var wcList []wordCount
	for k, v := range freq {
		wcList = append(wcList, wordCount{k, v})
	}
	sort.Slice(wcList, func(i, j int) bool {
		return wcList[i].count > wcList[j].count
	})
	fmt.Println("\nTop", top, "words:")
	for i := 0; i < top && i < len(wcList); i++ {
		fmt.Printf("%s: %d\n", wcList[i].word, wcList[i].count)
	}
}
