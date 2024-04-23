package main

import (
	"bufio"
	"fmt"
	"github.com/mainul35/go-with-abc/main/back-to-basics/structpractice/structs"
	"os"
	"strconv"
	"strings"
)

func main() {

	var people []structs.Person
	var reader = bufio.NewReader(os.Stdin)
	for i := 1; i <= 2; i++ {
		println("Enter name: ")
		var name, _ = reader.ReadString('\n')
		name = strings.TrimRight(name, "\r\n")
		println("Enter age: ")
		var age, _, _ = reader.ReadLine()
		ageInt, _ := strconv.Atoi(string(age))
		p := structs.Person{Name: name, Age: ageInt}
		people = append(people, p)
	}

	println("Printing all the names and ages available in the array")
	for i := 0; i < len(people); i++ {
		fmt.Print("Name: " + people[i].Name + ", Age: " + strconv.Itoa(people[i].Age))
		fmt.Println()
	}
}
