package main

import (
	"errors"
	"fmt"
	"time"
)

type person struct {
	firstName string
	lastName  string
	birth     string
	createdAt time.Time
}

func main() {
	firstName, _ := getUserInput("Enter first name: ")
	lastName, _ := getUserInput("Enter last name: ")
	birth, _ := getUserInput("Enter birth (MM/DD/YYYY): ")

	person, err := newPerson(firstName, lastName, birth)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(person.String())
	person.ClearName()
	fmt.Println(person.String())
}

func newPerson(firstName, lastName, birth string) (*person, error) {
	if firstName == "" || lastName == "" || birth == "" {
		return nil, errors.New("invalid name or birth")
	}

	return &person{
		firstName: firstName,
		lastName:  lastName,
		birth:     birth,
		createdAt: time.Now(),
	}, nil
}

func (p *person) String() string {
	return fmt.Sprintf("Name: %s %s\nBirth: %s\nCreated at: %s\n", p.firstName, p.lastName, p.birth, p.createdAt)
}

func (p *person) ClearName() {
	p.firstName = ""
	p.lastName = ""
}

func getUserInput(text string) (value string, err error) {
	fmt.Print(text)
	fmt.Scanln(&value)

	return
}
