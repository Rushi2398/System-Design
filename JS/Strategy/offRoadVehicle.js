import { Vehicle } from "./vehicle.js";
import { NormalDrive } from "./normalDrive.js";

export class OffRoadVehicle extends Vehicle {
  constructor() {
    super(new NormalDrive());
  }
}
