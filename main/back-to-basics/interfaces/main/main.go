package main

import (
	"github.com/mainul35/go-with-abc/main/back-to-basics/interfaces/models"
	"log"
)

func main() {
	car := models.Car{Brand: "Honda", Model: "CR-V", Make: "2002", TransmissionType: "4 Speed Automatic"}

	res := car.GetVehicleInfo()

	log.Println(res)
}
