package main

import "log"

func main() {
	myMap := make(map[string]string)
	myMap["name"] = "Syed Mainul Hasan"
	myMap["age"] = "30"
	myMap["nationality"] = "Bangladeshi"
	myMap["profession"] = "Software Engineer"

	for key, value := range myMap {
		log.Println(key, ":", value)
	}
}
