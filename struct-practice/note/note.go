package note

import (
	"errors"
	"fmt"
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
