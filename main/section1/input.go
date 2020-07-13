package main

import (
	"fmt"
	"os"
)

func main() {
	input := "yes"

	var result bool
	switch input {
	case "true", "yes", "1":
		result = true
	case "false", "no", "0":
		result = false
	default:
		fmt.Println("Error when converting to bool!")
		os.Exit(1)
	}

	fmt.Println("Your answer accepted as:", result)
}
