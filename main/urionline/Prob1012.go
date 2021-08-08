package main

import "fmt"

func main() {

	var A float64 = 0
	var B float64 = 0
	var C float64 = 0

	fmt.Scanf("%f", &A)
	fmt.Scanf("%f", &B)
	fmt.Scanf("%f", &C)

	fmt.Printf("TRIANGULO: %.3f\n", (.5 * A * C))
	fmt.Printf("CIRCULO: %.3f\n", (3.14159 * C * C))
	fmt.Printf("TRAPEZIO: %.3f\n", ((A+B)/2)*C)
	fmt.Printf("QUADRADO: %.3f\n", (B * B))
	fmt.Printf("RETANGULO: %.3f\n", (A * B))
}
