package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {
	if text == "" {
		return &Todo{}, errors.New("invalid text content")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (n *Todo) String() string {
	return fmt.Sprintf("Text: %s\n", n.Text)
}

func (n Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(json), 0644)
}
