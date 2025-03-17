package price

import (
	"fmt"
)

type PriceWithTaxJob struct {
	TaxRate        float64
	Prices         []float64
	PriceWithTaxes map[string]float64
}

func (p *PriceWithTaxJob) Process() {
	result := make(map[string]float64)
	for _, price := range p.Prices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + p.TaxRate)
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
