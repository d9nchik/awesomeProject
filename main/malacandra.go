package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	for count := 10; count > 0; count-- {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			fmt.Print("Something gone wrong")
			os.Exit(1)
		}
	}
	fmt.Println("Liftoff!")
}
