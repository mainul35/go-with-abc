package models

import "fmt"

type Car struct {
	Brand            string
	Model            string
	Make             string
	TransmissionType string
}

func (c *Car) GetVehicleInfo() string {
	s := "Brand: %s\nModel: %sMake: %s\nEngine Model: %s\nTransmission Type:%s"
	return fmt.Sprintf(s, c.Brand, c.Model, c.Make, c.GetEngineInfo(), c.TransmissionType)
}

func (c *Car) GetEngineInfo() string {
	return "K20A2"
}
