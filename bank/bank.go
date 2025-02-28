package main

import "fmt"

func main() {
	var balance float64 = 1000.0

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

		if choice == 1 {
			fmt.Println("Enter amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount")
				break
			}

			balance += amount
			fmt.Println("Your balance is:", balance)
		} else if choice == 2 {
			fmt.Println("Enter amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount")
				break
			}

			if amount > balance {
				fmt.Println("Insufficient balance")
				break
			}

			balance -= amount
			fmt.Println("Your balance is:", balance)
		} else if choice == 3 {
			fmt.Println("Your balance is:", balance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for using Go Bank!")
}
