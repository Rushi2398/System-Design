export class Vehicle {
  constructor(driveObj) {
    this.driveObj = driveObj;
  }
  driveVehicle() {
    this.driveObj.drive();
  }
}
