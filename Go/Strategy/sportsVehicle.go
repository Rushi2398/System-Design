package main

type SportsVehicle struct {
	Vehicle
}

func NewSportsVehicle() *SportsVehicle {
	return &SportsVehicle{
		Vehicle{driveBehaviour: SpecialDrive{}},
	}
}
