package fileop

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadBalanceFromFile(fileName string) (float64, error) {
	balanceByte, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("balance file not found")
	}

	balanceString := string(balanceByte)
	balance, err := strconv.ParseFloat(balanceString, 64)
	if err != nil {
		return 1000, errors.New("invalid balance")
	}

	return balance, nil
}

func WriteBalanceToFile(balance float64, fileName string) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceString), 0644)
}
