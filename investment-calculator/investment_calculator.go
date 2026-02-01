package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Years: ")
	fmt.Scan(&years)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	calculateFutureValues(investmentAmount, expectedReturnRate, years)
	fmt.Println("Real Rate Calculation:", calculateRealRate(expectedReturnRate, inflationRate))
	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Real Value:", futureRealValue)
}

func calculateFutureValues(principal float64, rate float64, years float64) (float64, float64) {
	futureValue := principal * math.Pow(1+rate/100, years)
	futureRealValue := futureValue / math.Pow(1+rate/100, years)

	return futureValue, futureRealValue
}

func calculateRealRate(nominalRate float64, inflationRate float64) float64 {
	return (1+nominalRate)/(1+inflationRate) - 1
}
