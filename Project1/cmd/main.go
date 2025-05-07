package main

import (
	"fmt"
	"github.com/Luffy-SunGod/ConcurrencyInGolang/Project1"
	"runtime"
	"sync"
	"time"
)

type ChannelCurrencyStruct struct {
	Code string
	Name string
}

func main() {
	numWorkers := runtime.NumCPU()
	//fmt.Println(numWorkers)
	ce := Project1.MyCurrencyExchange{
		Currencies: map[string]Project1.Currency{},
	}
	err := ce.FetchAllCurrencies()
	if err != nil {
		fmt.Println(err)
		return
	}
	startTime := time.Now()

	job := make(chan ChannelCurrencyStruct, 10)
	result := make(chan Project1.Currency, 10)

	wg := &sync.WaitGroup{}
	for i := 0; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(job, result, wg)
	}

	for key, val := range ce.Currencies {
		s := ChannelCurrencyStruct{
			Code: key,
			Name: val.Name,
		}
		job <- s
	}
	close(job)
	go func() {
		wg.Wait()
		close(result)
	}()

	finalResult := map[string]Project1.Currency{}
	fmt.Println("------Result------")
	for val := range result {
		c := Project1.Currency{
			Code:  val.Code,
			Name:  val.Name,
			Rates: val.Rates,
		}
		finalResult[val.Code] = c
		fmt.Printf("%s (%s): %d rates\n", val.Name, val.Code, len(val.Rates))
	}
	endTime := time.Now()
	fmt.Println("-------------------")
	fmt.Println("Time Taken: ", endTime.Sub(startTime))
}

func Worker(job chan ChannelCurrencyStruct, res chan Project1.Currency, group *sync.WaitGroup) {
	defer group.Done()

	for jobStruct := range job {
		response, _ := Project1.FetchCurrencyRates(jobStruct.Code)
		c := Project1.Currency{
			Code:  jobStruct.Code,
			Name:  jobStruct.Name,
			Rates: response,
		}
		res <- c
	}
}
