package main

import (
	"fmt"
	"strings"
)

func main() {
	message := strings.ToLower("Hola Estación Espacial Internacional")
	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
}
