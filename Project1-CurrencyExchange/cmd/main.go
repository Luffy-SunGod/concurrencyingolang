package main

import (
	"fmt"
	"github.com/Luffy-SunGod/ConcurrencyInGolang/Project1-CurrencyExchange"
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
	ce := Project1_CurrencyExchange.MyCurrencyExchange{
		Currencies: map[string]Project1_CurrencyExchange.Currency{},
	}
	err := ce.FetchAllCurrencies()
	if err != nil {
		fmt.Println(err)
		return
	}
	startTime := time.Now()

	job := make(chan ChannelCurrencyStruct, 10)
	result := make(chan Project1_CurrencyExchange.Currency, 10)

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

	finalResult := map[string]Project1_CurrencyExchange.Currency{}
	fmt.Println("------Result------")
	for val := range result {
		c := Project1_CurrencyExchange.Currency{
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

func Worker(job chan ChannelCurrencyStruct, res chan Project1_CurrencyExchange.Currency, group *sync.WaitGroup) {
	defer group.Done()

	for jobStruct := range job {
		response, _ := Project1_CurrencyExchange.FetchCurrencyRates(jobStruct.Code)
		c := Project1_CurrencyExchange.Currency{
			Code:  jobStruct.Code,
			Name:  jobStruct.Name,
			Rates: response,
		}
		res <- c
	}
}
