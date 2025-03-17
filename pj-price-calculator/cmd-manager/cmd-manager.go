package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() *CMDManager {
	return &CMDManager{}
}

func (cm *CMDManager) ReadFile() ([]string, error) {
	fmt.Println("Please enter the prices. Confirm with enter.")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scanln(&price)

		if price == "" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cm *CMDManager) WriteFile(data interface{}) error {
	fmt.Println(data)
	return nil
}
