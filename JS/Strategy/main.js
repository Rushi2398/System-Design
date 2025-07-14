import { SportsVehicle } from "./sportsVehicle.js";
import { OffRoadVehicle } from "./offRoadVehicle.js";
import { CityVehicle } from "./cityVehicle.js";

const sportsCar = new SportsVehicle();
const offRoadCar = new OffRoadVehicle();
const cityCar = new CityVehicle();

sportsCar.driveVehicle(); // Special drive capability
offRoadCar.driveVehicle(); // Normal drive capability
cityCar.driveVehicle(); // Normal drive capability
