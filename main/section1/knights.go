package main

import (
	"fmt"
)

type item string

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil {
		fmt.Println("This character hasn`t been created")
	} else {
		c.leftHand = i
		fmt.Println(c.name, "got", *i)
	}
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Println(c.name, "has given", *to.leftHand, "to", to.name)

}

func main() {
	arthur := character{name: "Arthur"}
	knight := character{
		name: "Knight",
	}
	sword := item("sword")
	arthur.pickup(&sword)
	arthur.give(&knight)
}
