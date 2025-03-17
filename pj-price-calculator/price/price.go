package price

import (
	"fmt"

	"example.com/pj-price-calculator/conversion"
	filemanager "example.com/pj-price-calculator/file-manager"
)

type PriceWithTaxJob struct {
	TaxRate        float64
	Prices         []float64
	PriceWithTaxes map[string]string
}

func (p *PriceWithTaxJob) LoadData() {
	lines, err := filemanager.ReadFile("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
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

	p.PriceWithTaxes = result
	filemanager.WriteFileJSON(fmt.Sprintf("result_%.2f.json", p.TaxRate*100), p)
}

func New(taxRate float64) *PriceWithTaxJob {
	return &PriceWithTaxJob{
		TaxRate:        taxRate,
		Prices:         []float64{10, 20, 30},
		PriceWithTaxes: make(map[string]string),
	}
}
