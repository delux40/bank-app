package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"project.com/bank/fileops"
)

const accBalanceFileName = "Balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accBalanceFileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		panic("Can`t continue, sorry.")
	}

	fmt.Println("Welcome to Go bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposite amount. Myst be greater then 0!")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! Your balance now is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accBalanceFileName)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount. Myst be greater then 0!")
				continue
			}
			if accountBalance < withdrawAmount {
				fmt.Print("Error. You need more money on the balance!")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Print("Balance updated! Your balance now is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accBalanceFileName)
		case 4:
			fmt.Println("Thanks for choosing our bank!")
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Choose correct option!")
			continue
		}
	}
}
