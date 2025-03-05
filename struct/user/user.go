package user

import (
	"errors"
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birth     string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	Person
}

func NewAdmin(email, password string) Admin {
	if email == "" || password == "" {
		return Admin{}
	}

	return Admin{
		email:    email,
		password: password,
		Person: Person{
			firstName: "Admin",
			lastName:  "",
			birth:     "--",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birth string) (*Person, error) {
	if firstName == "" || lastName == "" || birth == "" {
		return nil, errors.New("invalid name or birth")
	}

	return &Person{
		firstName: firstName,
		lastName:  lastName,
		birth:     birth,
		createdAt: time.Now(),
	}, nil
}

func (p *Person) String() string {
	return fmt.Sprintf("Name: %s %s\nBirth: %s\nCreated at: %s\n", p.firstName, p.lastName, p.birth, p.createdAt)
}

func (p *Person) ClearName() {
	p.firstName = ""
	p.lastName = ""
}
