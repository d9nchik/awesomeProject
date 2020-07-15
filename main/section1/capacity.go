package main

import "fmt"

type integerSlice []int32

func main() {
	var slice integerSlice

	for i := 0; i < 40; i++ {
		slice = append(slice, int32(i))
		slice.info()
	}
}

func (slice integerSlice) info() {
	fmt.Printf("Added new item: Length: %d; Capacity %d\n", len(slice), cap(slice))
}
