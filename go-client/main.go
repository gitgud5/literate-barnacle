package main

import (
	"fmt"
	"sync"

	"cryptobro.com/cool/api"
)

func main() {

	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go GetCurrencyData(currency, &wg)
	}
	wg.Wait()

}

func GetCurrencyData(currency string, wg *sync.WaitGroup) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	}
	wg.Done()
}
