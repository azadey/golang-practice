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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", note.Title, note.Content, note.CreatedAt.Format(time.RFC1123))
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_") + ".json"
	fileName = strings.ToLower(fileName)

	jsonContent, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0644)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid note data")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
