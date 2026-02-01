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

func (t Todo) Display() {
	fmt.Println(t.Text)
}

func (t Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(t) // convert Note to JSON

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Input cannot be empty")
	}
	return Todo{
		Text: text,
	}, nil
}
