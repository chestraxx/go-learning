package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, errR := getUserInput("Enter revenue: ")
	expense, errE := getUserInput("Enter expense: ")
	taxRate, errT := getUserInput("Enter tax rate: ")
	if errR != nil || errE != nil || errT != nil {
		fmt.Println(errR, errE, errT)
		return
	}

	ebt, profit, ratio := calcFinancial(revenue, expense, taxRate)

	formattedEBT := fmt.Sprintf("EBT: %.1f \n", ebt)
	formattedProfit := fmt.Sprintf("Profit: %.1f \n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.1f \n", ratio)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)
	storeResult(ebt, profit, ratio)
}

func storeResult(ebt, profit, ratio float64) {
	os.WriteFile("result.txt", []byte(fmt.Sprintf("EBT: %.1f \nProfit: %.1f \nRatio: %.1f \n", ebt, profit, ratio)), 0644)
}

func calcFinancial(revenue, expense, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func getUserInput(text string) (value float64, err error) {
	fmt.Print(text)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("invalid value")
	}

	return
}
