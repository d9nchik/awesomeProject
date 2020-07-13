package main

import (
	"fmt"
	"math/rand"
)

func main() {
	account := 0
	for account/100 < 20 {
		switch rand.Intn(3) {
		case 0:
			account += 5
			fmt.Println("Adding nickel: 0.05")
		case 1:
			account += 1
			fmt.Println("Adding dime: 0.1")
		case 2:
			account += 25
			fmt.Println("Adding quarters")
		}
		fmt.Println("Balance: $", float32(account)/100)
	}

	fmt.Println("Now you can buy a gift!")
	fmt.Println("Your balance is $", float32(account)/100.0)
}
