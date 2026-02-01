package main

import (
	"fmt"

	"github.com/PavelBLab/price-calculator/filemanager"
	"github.com/PavelBLab/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}

	doneChannels := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChannels[index] = make(chan bool)

		fm := filemanager.New("prices.txt", fmt.Sprintf("tax_included_prices_%.0f", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.New(fm, taxRate)

		go priceJob.ProcessPrices(doneChannels[index]) // run as goroutine

		//if err != nil {
		//	fmt.Printf("Error processing prices with tax rate %.2f: %v\n", taxRate, err)
		//	break
		//}

	}

	for _, doneChan := range doneChannels {
		<-doneChan
	}
}
