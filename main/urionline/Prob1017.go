package main

import "fmt"

func main() {

	var h, speed int64
	fmt.Scanf("%d", &h)
	fmt.Scanf("%d", &speed)
	var spentFuel float64
	spentFuel = float64(h*speed) / float64(12)
	fmt.Printf("%.3f\n", spentFuel)
}
