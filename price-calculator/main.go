package main

import (
	"fmt"

	"github.com/PavelBLab/price-calculator/filemanager"
	"github.com/PavelBLab/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}

	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChannels[index] = make(chan bool)
		errorChannels[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("tax_included_prices_%.0f", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.New(fm, taxRate)

		go priceJob.ProcessPrices(doneChannels[index], errorChannels[index]) // run as goroutine

		//if err != nil {
		//	fmt.Printf("Error processing prices with tax rate %.2f: %v\n", taxRate, err)
		//	break
		//}

	}

	for index := range taxRates {
		select {
		case err := <-errorChannels[index]:
			if err != nil {
				fmt.Printf("Error processing prices with tax rate %.2f: %v\n", taxRates[index]*100, err)
				break
			}
		case <-doneChannels[index]:
			fmt.Printf("Prices processed with tax rate %.2f\n", taxRates[index]*100)
		}
	}
}
