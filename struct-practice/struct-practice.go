package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/struct-practice/note"
)

func getUserInput(text string) (value string, err error) {
	fmt.Print(text)

	reader := bufio.NewReader(os.Stdin)

	value, err = reader.ReadString('\n')
	if value == "" {
		return "", errors.New("invalid value")
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

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
