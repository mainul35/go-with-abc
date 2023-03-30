package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	name string
	age  int
}

func main() {

	var people []Person
	var reader = bufio.NewReader(os.Stdin)
	for i := 1; i <= 2; i++ {
		println("Enter name: ")
		var name, _ = reader.ReadString('\n')
		name = strings.TrimRight(name, "\r\n")
		println("Enter age: ")
		var age, _, _ = reader.ReadLine()
		ageInt, _ := strconv.Atoi(string(age))
		p := Person{name: name, age: ageInt}
		people = append(people, p)
	}

	println("Printing all the names and ages available in the array")
	for i := 0; i < len(people); i++ {
		fmt.Print("Name: " + people[i].name + ", Age: " + strconv.Itoa(people[i].age))
		fmt.Println()
	}
}
