package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountFile)

	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("------------------------------------------------")
		// panic("Could not get account balance. Please ensure the balance file exists and is correctly formatted.")
	}

	fmt.Println("Welcome to Bank Go!")
	fmt.Println("Your account number is:", randomdata.Number(10000000, 99999999))
	fmt.Println("Your current balance is:", accountBalance)
	fmt.Println("------------------------------------------------")

	for {
		displayMenu()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is:", accountBalance)

		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than zero.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Deposit successful. New balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountFile)

		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Withdrawal amount must be greater than zero.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds. Your balance is:", accountBalance)
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Withdrawal successful. New balance is:", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountFile)
			}

		default:
			fmt.Println("Thank you for using Bank Go. Goodbye!")
			return
		}

		fmt.Println("Thanks for using our banking services!")
	}
}
