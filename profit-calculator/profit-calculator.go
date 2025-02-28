package main

import "fmt"

func main() {
	revenue := getUserInput("Enter revenue: ")
	expense := getUserInput("Enter expense: ")
	taxRate := getUserInput("Enter tax rate: ")

	ebt, profit, ratio := calcFinancial(revenue, expense, taxRate)

	formattedEBT := fmt.Sprintf("EBT: %.1f \n", ebt)
	formattedProfit := fmt.Sprintf("Profit: %.1f \n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.1f \n", ratio)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)
}

func calcFinancial(revenue, expense, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func getUserInput(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)
	return
}
