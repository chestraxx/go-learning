package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/struct-practice/note"
	"example.com/struct-practice/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	String() string
}

type outputtable interface {
	saver
	displayer
}

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

func getTodoData() (value string, err error) {
	return getUserInput("Todo text: ")
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saved failed")
		return err
	}

	fmt.Println("Saved successfully")
	return nil
}

func outputData(data outputtable) {
	fmt.Print(data.String())
	saveData(data)
}

func main() {
	todoText, err := getTodoData()
	if err != nil {
		fmt.Println(err)
		return
	}

	todoNote, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(todoNote)

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
	outputData(note)
}
