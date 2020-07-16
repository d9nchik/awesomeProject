package main

import "fmt"

type location struct {
	name      string
	lat, long float64
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func main() {
	landingSiteOnMars := []location{
		newLocation(coordinate{d: 14, m: 34, s: 6.2, h: 'S'}, coordinate{175, 28, 21.5, 'E'},
			"Columbia Memorial Station"),

		newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'},
			"Challendger Memorial Station"),

		newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'},
			"Bradbury Landing"),
		newLocation(coordinate{4, 30, 0, 'N'}, coordinate{135, 54, 0, 'E'},
			"Elysium Planitia"),
	}

	for _, mar := range landingSiteOnMars {
		fmt.Println(mar)
	}
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate, name string) location {
	return location{name: name, lat: lat.decimal(), long: long.decimal()}
}
