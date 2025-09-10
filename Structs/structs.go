package main

import (
	"fmt"

	"example.com/struc/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (YYYY-MM-DD): ")

	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	appUser.DisplayUserInfo()
	appUser.ClearScreen()

	appAdmin, err := user.NewAdmin("Admin", "User", "N/A", "admin@example.com", "securepassword")
	if err != nil {
		fmt.Println("Error creating admin:", err)
		return
	}

	appAdmin.DisplayUserInfo()
	appAdmin.ClearScreen()
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
