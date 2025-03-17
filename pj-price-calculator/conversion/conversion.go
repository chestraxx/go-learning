package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	prices := make([]float64, len(strings))

	for index, string := range strings {
		floatPrice, err := strconv.ParseFloat(string, 64)
		if err != nil {
			return nil, errors.New("error parsing price")
		}

		prices[index] = floatPrice
	}

	return prices, nil
}
