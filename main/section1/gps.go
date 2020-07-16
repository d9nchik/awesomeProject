package main

import (
	"fmt"
	"math"
)

type rover struct {
	gps
}

type gps struct {
	current, destination location
	world                world
}

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

func main() {
	mars := world{radius: 3389.5}
	start := location{name: "Bradbury Landing", lat: -4.5895, long: 137.4417}
	end := location{name: "Elysium Planita", lat: 4.5, long: 135.9}
	curiosity := rover{gps{world: mars, current: start, destination: end}}
	fmt.Println(curiosity.message())
}

func (l location) description() string {
	return fmt.Sprintf("Name: %v; latitude: %v, longitude: %v", l.name, l.lat, l.long)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))

	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*-clong)
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("Left: %.2f(km)", g.distance())
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}
