package main

import (
	"encoding/json"
	"fmt"
)

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\\\"%c", c.d, c.m, c.s, c.h)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{	"decimal": %v,	"dms": "%s",	"degrees": %v,	"minutes": %v,	"seconds": %.1f,
	"hemisphere": "%c"}`, c.decimal(), c.String(), c.d, c.m, c.s, c.h)), nil
}

func main() {
	elysium := coordinate{135, 54, 0, 'E'}
	data, err := json.Marshal(elysium)
	if err == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println(err.Error())
	}
}
