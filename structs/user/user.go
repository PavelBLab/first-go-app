package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	age       int
	createdAt time.Time
}

type Admin struct {
	User     // embedding
	email    string
	password string
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		//panic("All fields are required to create a new user")
		return nil, errors.New("missing required user fields")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func (u User) OutputUserDetails() {
	fmt.Println("User Details ===============")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt)
}
