package main

import "fmt"

func main() {
	var productNames [4]string
	productNames = [4]string{"apple"}
	productNames[2] = "banana"

	prices := [4]float64{10.99, 9.99, 44.99, 0.99}

	fmt.Println(prices, productNames)
	fmt.Println(prices[3])

	featuredPrices := prices[1:]
	highlightedPrices := prices[:1]
	fmt.Println(featuredPrices, highlightedPrices)
}
