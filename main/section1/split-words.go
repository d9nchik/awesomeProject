package main

import (
	"fmt"
	"regexp"
)

func main() {
	lines := []string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
	}

	first := make(chan string)
	second := make(chan string)
	go creatorOfChan(lines, first)
	go splitWords(first, second)
	printChanel(second)
}

func creatorOfChan(lines []string, c chan<- string) {
	for _, line := range lines {
		c <- line
	}
	close(c)
}

func splitWords(input <-chan string, output chan<- string) {
	reg := regexp.MustCompile("\\W+")
	for s := range input {
		for _, word := range reg.Split(s, -1) {
			output <- word
		}
	}
	close(output)
}

func printChanel(input <-chan string) {
	for s := range input {
		fmt.Println(s)
	}
}
