package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	createdAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("invalid title or content")
	}

	return &Note{
		Title:     title,
		Content:   content,
		createdAt: time.Now(),
	}, nil
}

func (n *Note) String() string {
	return fmt.Sprintf("Title: %s\nContent: %s\nCreated at: %s\n", n.Title, n.Content, n.createdAt)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName+".json", []byte(json), 0644)
}
