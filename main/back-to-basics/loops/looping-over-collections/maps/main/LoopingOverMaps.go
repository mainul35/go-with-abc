package main

import "log"

func main() {
	var fruits = []string{"Apple", "Orange", "Mango", "Grape"}

	for i, fruit := range fruits {
		log.Println(i, fruit)
	}
}
