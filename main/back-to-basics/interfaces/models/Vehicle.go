package models

type Vehicle interface {
	GetVehicleInfo() string
}

type Engine interface {
	GetEngineInfo() string
}
