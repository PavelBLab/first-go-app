package main

import (
	"errors"
	"fmt"
	"os"
)

const fileName = "profit-calculator/profit_calculations.txt"

func main() {
	var taxRate float64 = 25.8

	revenue, err1 := getUserInput("Enter Revenue: ")
	if err1 != nil {
		fmt.Println("Error reading revenue:", err1)
		panic(err1)
	}

	costs, err2 := getUserInput("Enter Costs: ")

	if err2 != nil {
		fmt.Println("Error reading costs:", err2)
		panic(err2)
	}

	ebt, tax, ebtMargin, netProfit := calculateProfit(revenue, costs, taxRate)

	writeCalculationsToFile(ebt, tax, ebtMargin, netProfit)

	fmt.Println("Earnings Before Tax:", ebt)
	fmt.Printf("Tax: %.2f\n", tax)
	fmt.Printf("Earnings Before Tax Margin (%%): %.2f\n", ebtMargin)
	fmt.Printf("Net Profit: %.2f\n", netProfit)
}

func getUserInput(prompt string) (float64, error) {
	var value float64
	fmt.Print(prompt)
	fmt.Scan(&value)

	if value <= 0 {
		return 0.0, errors.New("Input must be a positive number.")
	}

	return value, nil
}

func calculateProfit(revenue float64, costs float64, taxRate float64) (float64, float64, float64, float64) {
	ebt := revenue - costs
	tax := ebt * taxRate / 100
	var ebtMargin = ebt / revenue * 100
	netProfit := ebt - tax

	return ebt, tax, ebtMargin, netProfit
}

func writeCalculationsToFile(ebt float64, tax float64, ebtMargin float64, netProfit float64) {
	var data string = fmt.Sprintf("Earnings Before Tax: %.2f\nTax: %.2f\nEarnings Before Tax Margin (%%): %.2f\nNet Profit: %.2f\n", ebt, tax, ebtMargin, netProfit)
	var dataBytes []byte = []byte(data)
	os.WriteFile(fileName, dataBytes, 0644)
}
