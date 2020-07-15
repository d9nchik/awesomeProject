package main

import (
	"fmt"
)

type stringSlice []string

func main() {
	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	appendString(planets[3:6], "New ")
	fmt.Println(planets)
}

func appendString(str stringSlice, prefix string) {
	for i := range str {
		str[i] = prefix + str[i]
	}
}
