package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil)
}

func About(writer http.ResponseWriter, request *http.Request) {
	// TODO: Implement
}

func Home(writer http.ResponseWriter, request *http.Request) {
	// TODO: Implement
}

func Index(writer http.ResponseWriter, request *http.Request) {
	bytesWritten, err := fmt.Fprintf(writer, "Hello world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of bytes written: %d", bytesWritten))
}
