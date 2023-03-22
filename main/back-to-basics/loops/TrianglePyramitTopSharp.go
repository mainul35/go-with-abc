package main

import "fmt"

func main() {
	target := 5
	for i := 0; i <= target; i++ {
		var j int = 0
		for ; j < target-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < target-j; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
