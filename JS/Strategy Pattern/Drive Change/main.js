import { Vehicle } from "./vehicle.js";
import { SpecialDrive } from "./specialDrive.js";
import { NormalDrive } from "./normalDrive.js";

const sportsCar = new Vehicle(new SpecialDrive());
const offRoadCar = new Vehicle(new SpecialDrive());
const cityCar = new Vehicle(new NormalDrive());

sportsCar.driveVehicle(); // Special drive capability
offRoadCar.driveVehicle(); // Normal drive capability
cityCar.driveVehicle(); // Normal drive capability
