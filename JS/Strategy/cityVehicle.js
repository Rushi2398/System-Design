import { Vehicle } from "./vehicle.js";
import { NormalDrive } from "./normalDrive.js";

export class CityVehicle extends Vehicle {
  constructor() {
    super(new NormalDrive());
  }
}
