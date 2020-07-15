package main

import "fmt"

func main() {
	dwarfs := [5]string{
		"Ceras",
		"Pluto",
		"Humea",
		"Makemake",
		"Eris",
	}

	arrCopy := dwarfs

	arrCopy[0] = "oops"

	fmt.Println(dwarfs)
	fmt.Println(arrCopy)
}
