package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var intNum int = 0
	//_,_ = fmt.Scanf("%d", &intNum)
	//fmt.Printf("The value of intNum is %d\n", intNum)
	//
	//var int8Num int8 = 0
	//_,_ = fmt.Scanf("%d", &int8Num)
	//fmt.Printf("The value of int8Num is %d\n", int8Num)
	//
	//var floatNum float32 = 0
	//_,_ = fmt.Scanf("%f", &floatNum)
	//fmt.Printf("The value of floatNum is %.2f\n", floatNum)

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Enter your full name: ")
	//myName,_ := reader.ReadString('\n')
	//println(myName)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter you age:")
	input,_ := reader.ReadString('\n')
	myAge,_ := strconv.ParseFloat(strings.TrimSpace(input), 16)
	fmt.Println(myAge + 2)
}
