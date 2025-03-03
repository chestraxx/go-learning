package main

import (
	"fmt"

	"example.com/bank/fileop"
)

const accountBalanceFileName = "balance.txt"

func main() {
	var balance, err = fileop.ReadBalanceFromFile(accountBalanceFileName)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("=======")
		panic("Error reading balance from file")
	}

	fmt.Println("Welcome to Go Bank!")
	for {
		showOption()

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
			fileop.WriteBalanceToFile(balance, accountBalanceFileName)
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
			fileop.WriteBalanceToFile(balance, accountBalanceFileName)
			fmt.Println("Your balance is:", balance)

		case 3:
			fileop.WriteBalanceToFile(balance, accountBalanceFileName)
			fmt.Println("Your balance is:", balance)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using Go Bank!")
			return
		}
	}
}
