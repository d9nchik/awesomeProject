package main

import (
	"fmt"
)

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

func newLocation(lat, long coordinate, name string) location {
	return location{name: name, lat: lat.decimal(), long: long.decimal()}
}
