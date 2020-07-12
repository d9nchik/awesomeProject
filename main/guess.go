package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const userNumber = 50
	lowerBounds := 1
	upperBounds := 100
	for {
		computerGuess := rand.Intn(upperBounds-lowerBounds) + lowerBounds
		fmt.Println("Computer think you guess number ", computerGuess)
		if computerGuess == userNumber {
			fmt.Println("He is right!")
			break
		} else if computerGuess > userNumber {
			fmt.Println("Oops! Too big")
			upperBounds = computerGuess
		} else {
			fmt.Println("Oops! Too small")
			lowerBounds = computerGuess
		}
	}
}
