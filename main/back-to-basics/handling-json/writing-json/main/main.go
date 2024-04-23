package main

import (
	"encoding/json"
	"fmt"
	handlingJson "github.com/mainul35/go-with-abc/main/back-to-basics/handling-json"
	"log"
)

func main() {
	var employees []handlingJson.Employee

	var emp1 handlingJson.Employee
	emp1.FirstName = "Syed"
	emp1.LastName = "Hasan"
	emp1.Age = 30
	emp1.Profession = "Software Engineer"

	employees = append(employees, emp1)

	byteArrOutput, err := json.MarshalIndent(employees, "", "    ")

	if err != nil {
		log.Println("Error marshalling", err)
	}

	fmt.Println(string(byteArrOutput))
}
