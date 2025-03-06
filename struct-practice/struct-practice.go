package main

import (
	"errors"
	"fmt"

	"example.com/struct-practice/note"
)

func getUserInput(text string) (value string, err error) {
	fmt.Print(text)
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("invalid value")
	}

	return
}

func getNoteData() (title, content string, err error) {
	title, err = getUserInput("Note title: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	content, err = getUserInput("Note content: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return
}

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(note.String())
}
