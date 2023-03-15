package main

import "fmt"

func main() {
	fmt.Println(addAndMultiply(2, 3))
	fmt.Println(addAndMultiplyWithNamedVars(2, 3))
}

func addAndMultiply(x int, y int) (int, int) {
	return x + y, x * y
}

func addAndMultiplyWithNamedVars(x int, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	//return x, y
	return
}
