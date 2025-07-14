package main

type OffRoadVehicle struct {
	Vehicle
}

func NewOffRoadVehicle() *OffRoadVehicle {
	return &OffRoadVehicle{
		Vehicle{driveBehaviour: NormalDrive{}},
	}
}
