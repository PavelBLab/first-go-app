package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PavelBLab/structs-practice/note"
	"github.com/PavelBLab/structs-practice/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

//

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Failed to get todo data:", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Failed to get note data:", err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println("Failed to save todo:", err)
		return
	}

	fmt.Println("Todo saved successfully.")

	err = outputData(userNote)

	if err != nil {
		fmt.Println("Failed to save note:", err)
		return
	}

	fmt.Println("Note saved successfully.")
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}

	fmt.Println("Data saved successfully.")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(promptText string) string {
	fmt.Printf("%v ", promptText)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
