package main

func main() {
	sportsCar := &Vehicle{driveBehaviour: SpecialDrive{}}
	offRoadCar := &Vehicle{driveBehaviour: NormalDrive{}}
	cityCar := &Vehicle{driveBehaviour: NormalDrive{}}

	sportsCar.DriveVehicle()  // Output: Sports drive capability
	offRoadCar.DriveVehicle() // Output: Normal drive capability
	cityCar.DriveVehicle()    // Output: Normal drive capability
}
