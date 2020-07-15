package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

//sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return offset + s()
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func _() {
	regulator := kelvin(5)
	sensor := calibrate(realSensor, regulator)
	fmt.Println(sensor())
	regulator = 6
	fmt.Println(sensor())

	sensor = calibrate(fakeSensor, regulator)
	fmt.Println(sensor())
	fmt.Println(sensor())
	fmt.Println(sensor())
}
