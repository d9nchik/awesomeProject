package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

type location struct {
	name      string
	lat, long float64
}

type coordinate struct {
	d, m, s float64
	h       rune
}

var landingSiteOnMars = []location{
	newLocation(coordinate{d: 14, m: 34, s: 6.2, h: 'S'}, coordinate{175, 28, 21.5, 'E'},
		"Columbia Memorial Station"),

	newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'},
		"Challendger Memorial Station"),

	newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'},
		"Bradbury Landing"),
	newLocation(coordinate{4, 30, 0, 'N'}, coordinate{135, 54, 0, 'E'},
		"Elysium Planitia"),
}

func main() {
	mars := world{radius: 3389.5}
	for i := range landingSiteOnMars {
		for j := i + 1; j < len(landingSiteOnMars); j++ {
			fmt.Printf("Distance between %v and %v is %v\n", landingSiteOnMars[i].name,
				landingSiteOnMars[j].name, mars.distance(landingSiteOnMars[i], landingSiteOnMars[j]))
		}
	}
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))

	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*-clong)
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
