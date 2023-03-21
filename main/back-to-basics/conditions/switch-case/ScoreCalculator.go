package main

import "fmt"

func main() {

	switch result := calculateResult(10); {
	case (result < 39):
		fmt.Println("Failed")
	case result >= 40 && result <= 49:
		fmt.Println("Passed with a D")
	case result >= 50 && result <= 59:
		fmt.Println("Passed with a C")
	case result >= 60 && result <= 69:
		fmt.Println("Passed with a B")
	case result >= 70 && result <= 100:
		fmt.Println("Passed with a A")
	default:
		fmt.Println("Passed")
	}
}

func calculateResult(score int) int {
	return 20 + score
}
