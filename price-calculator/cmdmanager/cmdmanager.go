package cmdmanager

import "fmt"

type CmdManager struct {
}

func New() *CmdManager {
	return &CmdManager{}
}

func (cmd *CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Enter the prices of the products:")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scanln(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd *CmdManager) WriteResults(data interface{}) error {
	fmt.Println(data)
	return nil
}
