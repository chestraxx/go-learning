package main

import (
	"fmt"

	"example.com/pj-price-calculator/price"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceWithTaxJob := price.New(taxRate)
		priceWithTaxJob.Process()

		// result[taxRate] = priceWithTaxJob.PriceWithTaxes
	}

	fmt.Println(result)
}
