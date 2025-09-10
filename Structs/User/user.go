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
	createdAt time.Time
}

type Admin struct {
	User
	email    string
	password string
}

func (user *User) DisplayUserInfo() {
	fmt.Println(user.firstName, user.lastName, "was born on", user.birthdate)
}

func (u *User) ClearScreen() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(firstName, lastName, birthdate, email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required for admin")
	}

	return &Admin{
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "N/A",
			createdAt: time.Now(),
		},
		email:    email,
		password: password,
	}, nil
}
