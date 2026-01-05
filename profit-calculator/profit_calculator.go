package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var costs float64
	var taxRate float64 = 25.8

	revenue = getUserInput("Enter Revenue: ")
	costs = getUserInput("Enter Costs: ")

	ebt, tax, ebtMargin, netProfit := calculateProfit(revenue, costs, taxRate)

	fmt.Println("Earnings Before Tax:", ebt)
	fmt.Printf("Tax: %.2f\n", tax)
	fmt.Printf("Earnings Before Tax Margin (%%): %.2f\n", ebtMargin)
	fmt.Printf("Net Profit: %.2f\n", netProfit)
}

func getUserInput(prompt string) float64 {
	var value float64
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}

func calculateProfit(revenue float64, costs float64, taxRate float64) (float64, float64, float64, float64) {
	ebt := revenue - costs
	tax := ebt * taxRate / 100
	var ebtMargin = ebt / revenue * 100
	netProfit := ebt - tax

	return ebt, tax, ebtMargin, netProfit
}
