package Project1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Currency struct {
	Code  string
	Name  string
	Rates map[string]float64
}

type ResponseApi1Struct struct {
	Code map[string]float64
}

type MyCurrencyExchange struct {
	Currencies map[string]Currency
}

func (m *MyCurrencyExchange) FetchAllCurrencies() error {
	rsp, err := http.Get("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies.json")

	if err != nil {
		fmt.Println("response getting error!!")
		return err
	}
	defer rsp.Body.Close()

	res, err := io.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("while reading response!")
		return err
	}
	csMap := make(map[string]string)
	err = json.Unmarshal(res, &csMap)
	if err != nil {
		fmt.Println("error unmarsahlling error!!!")
		return err
	}
	i := 0
	for key, val := range csMap {
		if i > 10 {
			break
		}
		c := Currency{
			Code:  key,
			Name:  val,
			Rates: map[string]float64{},
		}
		m.Currencies[key] = c
		i++
	}
	return nil
}

func FetchCurrencyRates(currency string) (map[string]float64, error) {
	resp, err := http.Get(fmt.Sprintf("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/%s.json", currency))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resultMap := map[string]interface{}{}
	err = json.Unmarshal(res, &resultMap)
	rateMap := map[string]float64{}
	for key, val := range resultMap[currency].(map[string]interface{}) {
		rateMap[key] = val.(float64)
	}
	return rateMap, nil

}
