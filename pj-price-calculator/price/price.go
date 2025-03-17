package price

import (
	"fmt"

	"example.com/pj-price-calculator/conversion"
	iomanager "example.com/pj-price-calculator/io-manager"
)

type PriceWithTaxJob struct {
	IOManager      iomanager.IOManager `json:"-"`
	TaxRate        float64             `json:"tax_rate"`
	Prices         []float64           `json:"prices"`
	PriceWithTaxes map[string]string   `json:"prices_with_taxes"`
}

func (p *PriceWithTaxJob) LoadData() error {
	lines, err := p.IOManager.ReadFile()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}

	p.Prices = prices
	return nil
}

func (p *PriceWithTaxJob) Process() error {
	err := p.LoadData()
	if err != nil {
		fmt.Println(err)
		return err
	}

	result := make(map[string]string)
	for _, price := range p.Prices {
		priceWithTax := price * (1 + p.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", priceWithTax)
	}

	p.PriceWithTaxes = result
	return p.IOManager.WriteFile(p)
}

func New(io iomanager.IOManager, taxRate float64) *PriceWithTaxJob {
	return &PriceWithTaxJob{
		IOManager:      io,
		TaxRate:        taxRate,
		Prices:         []float64{10, 20, 30},
		PriceWithTaxes: make(map[string]string),
	}
}
