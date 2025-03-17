package price

import (
	"fmt"

	"example.com/pj-price-calculator/conversion"
	filemanager "example.com/pj-price-calculator/file-manager"
)

type PriceWithTaxJob struct {
	IOManager      *filemanager.FileManager `json:"-"`
	TaxRate        float64                  `json:"tax_rate"`
	Prices         []float64                `json:"prices"`
	PriceWithTaxes map[string]string        `json:"prices_with_taxes"`
}

func (p *PriceWithTaxJob) LoadData() {
	lines, err := p.IOManager.ReadFile()
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
	p.IOManager.WriteFile(p)
}

func New(fm *filemanager.FileManager, taxRate float64) *PriceWithTaxJob {
	return &PriceWithTaxJob{
		IOManager:      fm,
		TaxRate:        taxRate,
		Prices:         []float64{10, 20, 30},
		PriceWithTaxes: make(map[string]string),
	}
}
