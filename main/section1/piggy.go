package main

import (
	"fmt"
	"math/rand"
)

func main() {
	account := 0.0
	for account < 20 {
		switch rand.Intn(3) {
		case 0:
			account += 0.05
			fmt.Println("Adding nickel: 0.05")
		case 1:
			account += 0.1
			fmt.Println("Adding dime: 0.1")
		case 2:
			account += 0.25
			fmt.Println("Adding quarters")
		}
		fmt.Printf("Balance: $%.2f\n", account)
	}

	fmt.Println("Now you can buy a gift!")
	fmt.Printf("Your balance is $%.2f\n", account)
}
