package main

import "fmt"

// kelvinToCelsius convert K to C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// celsiusToFahrenheit convert C to F
func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// kelvinToFahrenheit convert K to F
func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	fmt.Println("Checking functions for correctness:")
	fmt.Println("0 kelvin is approximately equal -459.67")
	fmt.Println("Function result with parameter 0:", kelvinToFahrenheit(0))
}
