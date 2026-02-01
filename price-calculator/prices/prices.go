package prices

import (
	"fmt"

	"github.com/PavelBLab/price-calculator/conversion"
	"github.com/PavelBLab/price-calculator/iomanager"
)

type TaxInclPriceJob struct {
	IOManager         iomanager.IOManager
	taxRate           float64
	inputPrices       []float64
	taxIncludedPrices map[string][]string
}

func New(iom iomanager.IOManager, taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		IOManager:         iom,
		taxRate:           taxRate,
		taxIncludedPrices: make(map[string][]string),
	}
}

type TaxInclPriceJobJson struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string][]string `json:"tax_included_prices"`
}

func (job *TaxInclPriceJob) ToJson() TaxInclPriceJobJson {
	return TaxInclPriceJobJson{
		TaxRate:           job.taxRate,
		InputPrices:       job.inputPrices,
		TaxIncludedPrices: job.taxIncludedPrices,
	}
}

func (job *TaxInclPriceJob) loadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println("Error reading prices:", err)
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("Error converting prices to float:", err)
		return err
	}

	job.inputPrices = prices

	return nil
}

func (job *TaxInclPriceJob) ProcessPrices(doneChan chan bool) {
	job.loadData()

	//if err != nil {
	//	fmt.Println("Error loading prices:", err)
	//	return err
	//}

	result := make(map[string][]string)
	taxRateKey := fmt.Sprintf("%.2f", job.taxRate)

	for _, price := range job.inputPrices {
		priceInclTaxFormatted := fmt.Sprintf("%.2f", price*(1+job.taxRate))
		result[taxRateKey] = append(result[taxRateKey], priceInclTaxFormatted)
	}

	job.taxIncludedPrices = result

	dto := job.ToJson()

	job.IOManager.WriteResults(dto)

	//if err != nil {
	//	fmt.Println("Error writing results:", err)
	//	return err
	//}

	doneChan <- true
}
