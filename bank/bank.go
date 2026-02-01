package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/PavelBLab/bank/fileops"
)

const accountBalanceFileName = "bank/balance.txt"

func main() {
	var accountBalance, err = fileops.GetFLoatFromFile(accountBalanceFileName)

	if err != nil {
		fmt.Println("ERROR")
		panic(err)
	}

	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOption()

		var choice int
		fmt.Print("Please select an option: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your balance is $%.2f.\n", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFileName)
			fmt.Printf("You have deposited $%.2f. New balance is $%.2f.\n", depositAmount, accountBalance)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds.")
				continue
			}

			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdrawal amount.")
				continue
			}

			accountBalance -= withdrawAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFileName)
			fmt.Printf("You have withdrawn $%.2f. New balance is $%.2f.\n", withdrawAmount, accountBalance)

		} else if choice == 4 {
			fmt.Println("Exiting the banking system.")
			break
		} else {
			fmt.Println("Invalid option selected.")
			break
		}
	}

	fmt.Println("Thank you for banking with us. Goodbye!")
}
