package main

import "log"

func main() {
	val := "Some value to be replaced"
	log.Println(val)
	passedByReference(&val)
	log.Println(val)
}

func passedByReference(val *string) {
	*val = "Value changed in the existing reference"
}
