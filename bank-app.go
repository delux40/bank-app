package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check the balance")
		fmt.Println("2. Deposite money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
