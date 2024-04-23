package main

import (
	"encoding/json"
	"fmt"
	handlingJson "github.com/mainul35/go-with-abc/main/back-to-basics/handling-json"
	readingJson "github.com/mainul35/go-with-abc/main/back-to-basics/handling-json/reading-json"
	"log"
)

func main() {
	var unmarshalled []handlingJson.Employee
	err := json.Unmarshal([]byte(readingJson.GetJsonSchema()), &unmarshalled)

	if err != nil {
		log.Fatalln("Error unmarshalling JSON", err)
	}

	marshalled, err := json.MarshalIndent(unmarshalled, "", "    ")

	if err != nil {
		log.Fatalln("Error marshalling object to json", err)
	}

	fmt.Println(string(marshalled))
}
