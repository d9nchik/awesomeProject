package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

func main() {
	_, err := url.Parse("https://a b.com/")
	if err != nil {
		if errs, ok := err.(*url.Error); ok {
			bytes, err := json.Marshal(errs)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(string(bytes))
		}
		//fmt.Printf("%#v", err)
		fmt.Println(err)
		os.Exit(1)
	}
}
