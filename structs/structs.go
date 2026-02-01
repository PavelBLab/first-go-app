package main

import (
	"fmt"
	"github.com/PavelBLab/structs/user"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "Max"
	name.log()

	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthdate := getUserData("Enter your birthdate (<MM/DD/YYYY>): ")

	var appUser *user.User
	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	admin := user.NewAdmin("test@gmail.com", "test1234")
	admin.OutputUserDetails()
	admin.ClearUserName()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	var input string
	fmt.Scanln(&input)

	return input
}
