package main

import "fmt"

func main() {
	cipherText := "HIDUDES"
	keyword := "GOLANG"

	var output string
	keywordLen := len(keyword)
	for i := range cipherText {
		output += string((cipherText[i]+keyword[i%keywordLen])%('Z'-'A'+1) + 'A')
	}
	fmt.Println(output)
}
