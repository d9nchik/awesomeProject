package main

import (
	"fmt"
	"math/rand"
	"time"
)

type animal interface {
	fmt.Stringer
	move() string
	eat() string
}

type rabbit struct{}

func (r rabbit) String() string {
	return "rabbit"
}

func (r rabbit) move() string {
	return "jump"
}

func (r rabbit) eat() string {
	foods := []string{
		"carrot",
		"cabbage",
	}
	return foods[rand.Intn(len(foods))]
}

type horse struct{}

func (h horse) String() string {
	return "horse"
}

func (h horse) move() string {
	return "riding"
}

func (h horse) eat() string {
	foods := []string{
		"oat",
		"cabbage",
		"hay",
	}
	return foods[rand.Intn(len(foods))]
}

type pig struct{}

func (p pig) String() string {
	return "pig"
}

func (p pig) move() string {
	return "pick"
}

func (p pig) eat() string {
	foods := []string{
		"carrot",
		"cucumber",
		"acorn",
	}
	return foods[rand.Intn(len(foods))]
}

type sheep struct{}

func (s sheep) String() string {
	return "sheep"
}

func (s sheep) move() string {
	return "eating grass"
}

func (s sheep) eat() string {
	foods := []string{
		"pumpkin",
		"raspberry",
		"potato",
	}
	return foods[rand.Intn(len(foods))]
}

type chicken struct{}

func (c chicken) String() string {
	return "chicken"
}

func (c chicken) move() string {
	return "peck"
}

func (c chicken) eat() string {
	foods := []string{
		"seeds",
		"wheat",
		"water",
	}
	return foods[rand.Intn(len(foods))]
}

func main() {
	animals := []animal{rabbit{}, horse{}, pig{}, sheep{}, chicken{}}
	for i := 0; i < 72; i++ {
		currentTime := i % 24
		fmt.Println("Current time is ", currentTime)
		if currentTime < 6 || currentTime > 18 {
			fmt.Println("Now is night!")
		} else {
			currentAnimal := animals[rand.Intn(len(animals))]
			switch rand.Intn(2) {
			case 0:
				printMove(currentAnimal)
			case 1:
				printEat(currentAnimal)

			}
		}

		time.Sleep(time.Second)
	}
}

func printMove(a animal) {
	fmt.Printf("%v is doing: %v\n", a.String(), a.move())
}

func printEat(a animal) {
	fmt.Printf("%v is eating: %v\n", a.String(), a.eat())
}
