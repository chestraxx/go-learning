package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var years float64
	var interestRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter interest rate: ")
	fmt.Scan(&interestRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+interestRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println("Investment amount:", investmentAmount, "Future value:", futureValue)
	fmt.Println("Investment amount:", investmentAmount, "Future real value:", futureRealValue)
}
