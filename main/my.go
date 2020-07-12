package main

import "fmt"
import "math/rand"

func main() {
	fmt.Printf("My weight on surface of Mars is %v kg, and I would be %v years old.\n",
		73*0.3783, 18*365/687)

	const lightSpeed = 100_800 * 24

	choice := rand.Intn(10) + 1
	fmt.Println("I guess a number: ", choice)

	choice = rand.Intn(10) + 1
	fmt.Println("I guess another number: ", choice)

	distance := 56_000_000 + rand.Intn((401-56)*1_000_000+1)
	fmt.Print("Your distance: ", distance)
}
