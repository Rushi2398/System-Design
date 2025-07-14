package main

type CityVehicle struct {
	Vehicle
}

func NewCityVehicle() *CityVehicle {
	return &CityVehicle{
		Vehicle{driveBehaviour: NormalDrive{}},
	}
}
