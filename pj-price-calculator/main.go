package main

import (
	"fmt"

	// cmdmanager "example.com/pj-price-calculator/cmd-manager"
	filemanager "example.com/pj-price-calculator/file-manager"
	"example.com/pj-price-calculator/price"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", "result_"+fmt.Sprintf("%.2f", taxRate*100)+".json")
		// cmd := cmdmanager.New()

		priceWithTaxJob := price.New(fm, taxRate)
		err := priceWithTaxJob.Process()
		if err != nil {
			fmt.Println("Process failed")
			fmt.Println(err)
		}
	}
}
