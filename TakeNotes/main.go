package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/take-note/note"
	"example.com/take-note/todo"
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

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")
	printSomething(todo)

	if err != nil {
		fmt.Println("Failed to create todo:", err)
		return
	}

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println("Failed to create note:", err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(note)
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		println("Integer : ", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		println("Float : ", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		println("String : ", stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed to save data:", err)
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

func getUserInput(prompt string) string {
	fmt.Printf("%v: ", prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text := strings.TrimSuffix(input, "\n")
	text = strings.TrimSuffix(text, "\r") // For Windows compatibility

	return text
}
