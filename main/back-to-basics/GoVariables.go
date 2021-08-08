package main

import "fmt"

func main () {
	fmt.Println("Hello world")

	var myint int8
	for i:= 0; i< 129; i++ {
		myint += 1
	}
	fmt.Println(myint)

	var maxFloat float32
	maxFloat = 16777216
	fmt.Println(maxFloat == maxFloat + 10)
	fmt.Println(maxFloat + 10)
	fmt.Println(int(maxFloat))
}
