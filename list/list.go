package main

import "fmt"

func main() {
	prices := []float64{10.99, 9.99}
	prices = append(prices, 44.99, 0.99)

	fmt.Println(prices)
}

// func main() {
// 	productNames := [4]string{"apple"}
// 	productNames[2] = "banana"

// 	prices := [4]float64{10.99, 9.99, 44.99, 0.99}

// 	fmt.Println(prices, productNames)
// 	fmt.Println(prices[3])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := prices[:1]
// 	fmt.Println(prices, len(prices), cap(prices))
// 	fmt.Println(featuredPrices, len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(highlightedPrices, len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices, len(highlightedPrices), cap(highlightedPrices))
// }
