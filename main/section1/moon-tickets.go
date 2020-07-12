package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var spaceProvider = []string{
		"Space Adventures",
		"SpaceX",
		"Virgin Galactic",
	}

	const distance = 62_100_000 //km
	const secondsInDays = 3600 * 24

	fmt.Println("Spaceline	 Days Trip type  Price")
	for i := 0; i < 38; i++ {
		fmt.Print("=")
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		price := rand.Intn(15) + 36
		speed := price - 20 //km/s
		days := distance / speed / (secondsInDays)

		typeTrip := "One-way"
		if rand.Intn(2) == 0 {
			price *= 2
			days *= 2
			typeTrip = "Round-trip"
		}

		provider := spaceProvider[rand.Intn(len(spaceProvider))]

		fmt.Printf("%-18s %2d %-10s $%4d\n", provider, days, typeTrip, price)
	}
}
