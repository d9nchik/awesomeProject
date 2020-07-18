package main

import "fmt"

func main() {
	words := []string{
		"go",
		"gopher",
		"my", //duplicated word
		"my",
		"concurrency",
	}
	first := make(chan string)
	second := make(chan string)
	go creator(words, first)
	go removeDuplicate(first, second)
	printer(second)
}

func creator(words []string, c chan<- string) {
	for _, word := range words {
		c <- word
	}
	close(c)
}

func removeDuplicate(input <-chan string, output chan<- string) {
	previousWord := ""
	for s := range input {
		if previousWord != s {
			output <- s
			previousWord = s
		}
	}
	close(output)
}

func printer(input <-chan string) {
	for s := range input {
		fmt.Println(s)
	}
}
