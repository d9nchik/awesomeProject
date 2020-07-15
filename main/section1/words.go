package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "As far as eye could reach he saw nothing but the stems of the great" +
		"\nplants about him receding in the violet shade, and far overhead the" +
		"\nmultiple transparency of huge leaves filtering the sunshine to the" +
		"\nsolemn splendour of twilight in which he walked. Whenever he felt" +
		"\nable he ran again; the ground continued soft and springy, covered with" +
		"\nthe same resilient weed which was the first thing his hands had" +
		"\ntouched in Malacandra. Once or twice a small red creature scuttled" +
		"\nacross his path, but otherwise there seemed to be no life stirring in the" +
		"\nwood; nothing to fear—except the fact of wandering unprovisioned" +
		"\nand alone in a forest of unknown vegetation thousands or millions of" +
		"\nmiles beyond the reach or knowledge of man."
	stringMap := wordToMap(text)

	for k, v := range stringMap {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func wordToMap(text string) map[string]int32 {
	output := map[string]int32{}
	for _, s := range regexp.MustCompile("\\W+").Split(strings.ToLower(text), -1) {
		output[s]++
	}
	return output
}
