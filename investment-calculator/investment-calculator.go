package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var interestRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter interest rate: ")
	fmt.Scan(&interestRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calcFutureValue(investmentAmount, interestRate, years)

	fmt.Println("Investment amount:", investmentAmount, "Future value:", futureValue)
	fmt.Println("Investment amount:", investmentAmount, "Future real value:", futureRealValue)
}

func calcFutureValue(investmentAmount, interestRate, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow((1+interestRate/100), years)
	futureRealValue = futureValue / math.Pow((1+inflationRate/100), years)

	return
}
