package main

import (
	"github.com/mainul35/go-with-abc/main/back-to-basics/receivers"
	"log"
)

func main() {
	person := receivers.Person{}

	person.SetName("Syed Mainul Hasan")
	person.SetAge(30)

	log.Println(person.Name(), person.Age())
}
