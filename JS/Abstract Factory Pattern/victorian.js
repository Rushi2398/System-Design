import { Chair, Sofa, FurnitureFactory } from "./interface.js";

class VictorianChair extends Chair {
  sit() {
    console.log("Sitting on a Victorian chair");
  }
}

class VictorianSofa extends Sofa {
  lieDown() {
    console.log("Lying on a Victorian sofa");
  }
}

export class VictorianFurnitureFactory extends FurnitureFactory {
  createChair() {
    return new VictorianChair();
  }
  createSofa() {
    return new VictorianSofa();
  }
}
