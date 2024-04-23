package main

import (
	"fmt"
)

func main() {
	aString := "This is a string value"

	for i, l := range aString {
		fmt.Println(i, string(l))
	}
}
