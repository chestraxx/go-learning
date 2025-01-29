package main

import "fmt"

func main() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expense: ")
	fmt.Scan(&expense)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("EBT: %.1f \n", ebt)
	fmt.Printf("Profit: %.1f \n", profit)
	fmt.Printf("Ratio: %.1f \n", ratio)

}
