package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [height][width]bool

func main() {
	a := NewUniverse()
	a.seed()
	b := NewUniverse()
	for i := 0; i < 100; i++ {
		a.Show()
		Step(a, b)
		a, b = b, a
		time.Sleep(3_000 * time.Second)
		fmt.Print("\033[H")
	}
}

func (u *Universe) Show() {
	for i := range u {
		for _, element := range u[i] {
			if element {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func NewUniverse() *Universe {
	return new(Universe)
}

func (u *Universe) seed() {
	for i := 0; i < width*height*0.25; i++ {
		u[rand.Intn(height)][rand.Intn(width)] = true
	}
}

func (u *Universe) Alive(x, y int) bool {
	return u[(x+height)%height][(y+width)%width]
}

func (u *Universe) Neighbours(x, y int) int {
	neighbours := [...][2]int{
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
	}

	counter := 0
	for _, neighbour := range neighbours {
		if u.Alive(x+neighbour[0], y+neighbour[1]) {
			counter++
		}
	}
	return counter
}

func (u *Universe) Next(x, y int) bool {
	neighbours := u.Neighbours(x, y)
	return neighbours == 3 || (u.Alive(x, y) && neighbours == 2)
}

func Step(a, b *Universe) {
	for i := range a {
		for j := range a[i] {
			b[i][j] = a.Next(i, j)
		}
	}
}
