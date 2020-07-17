package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.x++
}

func (t *turtle) down() {
	t.x--
}

func (t *turtle) right() {
	t.y++
}

func (t *turtle) left() {
	t.y--
}

func main() {
	t := turtle{
		x: 0,
		y: 0,
	}

	t.up()
	t.left()
	t.left()
	t.down()
	t.down()
	t.right()

	fmt.Println(t)
}
