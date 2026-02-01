package main

import (
	"fmt"

	"github.com/PavelBLab/structs-practice/note"
	"github.com/PavelBLab/structs-practice/todo"
)

func main() {

	result := add1(5, 10)
	fmt.Println("add1 result:", result)

}

// Or any type example with generics
func add2[T int | float64 | string](a, b T) T {
	return a + b
}

func add1(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}
	return nil
}

// empty interface example and type assertion/switch
func printSomething1(value interface{}) {
	switch value.(type) { // type switch
	case note.Note:
		fmt.Println("This is a note")
	case todo.Todo:
		fmt.Println("This is a todo")
	default:
		fmt.Println("Unknown type")
	}
}

func printSomething2(value interface{}) {
	intValue, ok := value.(int) // type assertion
	if ok {
		fmt.Println("Integer value:", intValue)
	}

	stringValue, ok := value.(string)
	if ok {
		fmt.Println("String value:", stringValue)
	}
}
