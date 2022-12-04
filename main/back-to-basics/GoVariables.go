package main

import "fmt"

func main() {
	var myName = "Syed Mainul Hasan"
	fmt.Println("My name is", myName)

	var name string
	name = "Kathy"
	fmt.Println("name=", name)

	var sum int
	fmt.Println("The sum is", sum)

	part1, _ := 3, 0
	part2, _ := 2, 1
	sum = part1 + part2
	fmt.Println("sum = ", sum)

	var (
		lessonName = "variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName = ", lessonName)
	fmt.Println("lessonType = ", lessonType)

	word1, word2, _ := "Hello", "world", "!"
	fmt.Println(word1, word2)
}
