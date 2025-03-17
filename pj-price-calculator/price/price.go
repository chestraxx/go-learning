package price

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PriceWithTaxJob struct {
	TaxRate        float64
	Prices         []float64
	PriceWithTaxes map[string]float64
}

func (p *PriceWithTaxJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	prices := make([]float64, len(lines))
	for index, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error parsing price", err)
			return
		}

		prices[index] = floatPrice
	}

	p.Prices = prices
}

func (p *PriceWithTaxJob) Process() {
	p.LoadData()

	result := make(map[string]string)
	for _, price := range p.Prices {
		priceWithTax := price * (1 + p.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", priceWithTax)
	}

	fmt.Println(result)
	// p.PriceWithTaxes = result
}

func New(taxRate float64) *PriceWithTaxJob {
	return &PriceWithTaxJob{
		TaxRate:        taxRate,
		Prices:         []float64{10, 20, 30},
		PriceWithTaxes: make(map[string]float64),
	}
}
