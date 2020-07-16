package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude" xml:"longitude"`
	}

	curriosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curriosity)
	if err != nil {
		fmt.Println("Something gone wrong")
	} else {
		fmt.Println(string(bytes))
	}

}
