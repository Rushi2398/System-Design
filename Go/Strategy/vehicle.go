package main

type Vehicle struct {
	driveBehaviour DriveBehaviour
}

func (v *Vehicle) DriveVehicle() {
	v.driveBehaviour.Drive()
}
