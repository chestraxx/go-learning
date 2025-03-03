package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFileName = "balance.txt"

func readBalanceFromFile() float64 {
	balanceByte, _ := os.ReadFile(accountBalanceFileName)
	balanceString := string(balanceByte)

	balance, _ := strconv.ParseFloat(balanceString, 64)

	return balance
}

func writeBalanceToFile(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFileName, []byte(balanceString), 0644)
}

func main() {
	var balance float64 = readBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("Please select an option: ")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Balance")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			balance += amount
			writeBalanceToFile(balance)
			fmt.Println("Your balance is:", balance)

		case 2:
			fmt.Println("Enter amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			if amount > balance {
				fmt.Println("Insufficient balance")
				continue
			}

			balance -= amount
			writeBalanceToFile(balance)
			fmt.Println("Your balance is:", balance)

		case 3:
			writeBalanceToFile(balance)
			fmt.Println("Your balance is:", balance)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using Go Bank!")
			return
		}
	}

}
