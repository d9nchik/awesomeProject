package main

import "fmt"

func main() {
	var red uint8 = 255
	red++
	fmt.Println("red after incrementing: ", red)

	var number int8 = 127
	number++
	fmt.Println("number after incrementing", number)

	red--
	fmt.Println("red after decrementing: ", red)
	number--
	fmt.Println("number after decrementing", number)

}
