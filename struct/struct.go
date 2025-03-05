package main

import (
	"fmt"

	"example.com/struct/user"
)

func main() {
	firstName, _ := getUserInput("Enter first name: ")
	lastName, _ := getUserInput("Enter last name: ")
	birth, _ := getUserInput("Enter birth (MM/DD/YYYY): ")

	person, err := user.New(firstName, lastName, birth)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(person.String())
	person.ClearName()
	fmt.Println(person.String())
}

func getUserInput(text string) (value string, err error) {
	fmt.Print(text)
	fmt.Scanln(&value)

	return
}
