package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	marker := make(chan bool)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, marker)
	}

	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case <-marker:
		case <-timeout:
			fmt.Println("I`ve waited to much")
			return
		}
	}
}

func sleepyGopher(number int, input chan<- bool) {
	defer func() { input <- true }()
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("... snore...", number)
}
