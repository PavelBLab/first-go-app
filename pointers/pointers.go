package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", agePointer)  // memory address
	fmt.Println("Age:", *agePointer) // dereferencing

	getAdultYears(agePointer)
	fmt.Println("Age:", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
