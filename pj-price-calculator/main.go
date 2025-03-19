package main

import (
	"fmt"

	// cmdmanager "example.com/pj-price-calculator/cmd-manager"
	filemanager "example.com/pj-price-calculator/file-manager"
	"example.com/pj-price-calculator/price"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", "result_"+fmt.Sprintf("%.2f", taxRate*100)+".json")
		// cmd := cmdmanager.New()

		priceWithTaxJob := price.New(fm, taxRate)
		go priceWithTaxJob.Process(doneChans[index], errChans[index])
		// if err != nil {
		// 	fmt.Println("Process failed")
		// 	fmt.Println(err)
		// }
	}

	for index, _ := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println("Process", index, "failed")
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Process", index, "successfully")
		}
	}
}
