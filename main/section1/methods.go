package main

import "fmt"

type fahrenheit float64

type celsius float64

//celsius convert from kelvin to celsius
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

//fahrenheit convert from celsius to fahrenheit
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9/5 + 32)
}

//fahrenheit convert from kelvin to fahrenheit
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

//celsius convert from fahrenheit to celsius
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

//kelvin convert from celsius to kelvin
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

//kelvin convert from fahrenheit to kelvin
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}
func main() {
	fmt.Println("Checking functions for correctness:")
	fmt.Println("0 kelvin is approximately equal -459.67")
	fmt.Println("Function result with parameter 0:", kelvin(0).fahrenheit())
}
