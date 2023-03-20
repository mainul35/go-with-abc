package main
import "fmt"

func main() {

    switch result := calculateResult(10); {
        case (result < 39) : 
            fmt.Println("Failed")
        default: 
            fmt.Println("Passed")
    }
}

func calculateResult(score int) int {
    return 20 + score
}