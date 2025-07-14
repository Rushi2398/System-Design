package main

func main() {
	sportsCar := NewSportsVehicle()
	offRoadCar := NewOffRoadVehicle()
	cityCar := NewCityVehicle()

	sportsCar.DriveVehicle()  // Output: Sports drive capability
	offRoadCar.DriveVehicle() // Output: Normal drive capability
	cityCar.DriveVehicle()    // Output: Normal drive capability
}
