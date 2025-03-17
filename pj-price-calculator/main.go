package main

import (
	"fmt"

	filemanager "example.com/pj-price-calculator/file-manager"
	"example.com/pj-price-calculator/price"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("input.txt", "result_"+fmt.Sprintf("%.2f", taxRate*100)+".json")
		priceWithTaxJob := price.New(fm, taxRate)
		priceWithTaxJob.Process()
	}
}
