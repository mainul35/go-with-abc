package main

import (
	"github.com/mainul35/go-with-abc/main/back-to-basics/channels/util"
	"log"
)

const limit = 100

func calculateValue(intChan chan int) {
	randomNumber := util.CalculateRandom(limit)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	result := <-intChan

	log.Println(result)
}
