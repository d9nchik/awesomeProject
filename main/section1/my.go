package main

import "fmt"

func main() {
	account := 0.0
	for i := 0; i < 11; i++ {
		account += 0.1
	}
	fmt.Println(account)
}
