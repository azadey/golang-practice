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

func (todo Todo) Display() {
	fmt.Printf("Text: %s\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonContent, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0644)
}

func New(text string) (Todo, error) {

	if text == "" {
		return Todo{}, errors.New("Invalid todo data")
	}

	return Todo{
		Text: text,
	}, nil
}
