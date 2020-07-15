package main

import (
	"fmt"
	"math/big"
)

func _() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	daysPerYear := big.NewInt(365)

	distanceToCanisMajorDwarf := new(big.Int)
	distanceToCanisMajorDwarf.SetString("236000000000000000", 10)
	fmt.Println("Canis Major Dwarf is", distanceToCanisMajorDwarf, "km away.")
	seconds := new(big.Int)
	seconds.Div(distanceToCanisMajorDwarf, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	years := new(big.Int)
	years.Div(days, daysPerYear)
	fmt.Println("That is", years, "years of travel at light speed.")
}
