import { DriveInterface } from "./driveInterface.js";

export class NormalDrive extends DriveInterface {
  drive() {
    console.log("Normal drive capability");
  }
}
