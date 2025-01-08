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

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)

}
