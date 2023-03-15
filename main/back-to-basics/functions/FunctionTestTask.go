//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:

//* Call every function at least once

package main

import "fmt"

func main() {
	greet("Syed Mainul Hasan")
	printAndReturn()
	fmt.Println(addThreeNums(2, 3, 4))
	fmt.Println(getReturnedNum())
	firstNum, _ := returnTwoNums()
	fmt.Println(firstNum)
	addThreeNumbersTogetherUsingAnyCombinationOfExixtingFunctions()

}

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hello", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func printAndReturn() string {
	message := "Hello Syed"
	fmt.Println(message)
	return message
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addThreeNums(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// * Write a function that returns any number
func getReturnedNum() int {
	return 100
}

// * Write a function that returns any two numbers
func returnTwoNums() (int, int) {
	return 10, 20
}

// * Add three numbers together using any combination of the existing functions.
//   - Print the result
func addThreeNumbersTogetherUsingAnyCombinationOfExixtingFunctions() {
	firstNum := addThreeNums(5, 10, 15)
	secondNum := getReturnedNum()
	thirdNum, _ := returnTwoNums()
	fmt.Println(firstNum + secondNum + thirdNum)
}
