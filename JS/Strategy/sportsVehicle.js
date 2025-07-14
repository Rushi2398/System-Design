import { Vehicle } from "./vehicle.js";
import { SpecialDrive } from "./specialDrive.js";

export class SportsVehicle extends Vehicle {
  constructor() {
    super(new SpecialDrive());
  }
}
