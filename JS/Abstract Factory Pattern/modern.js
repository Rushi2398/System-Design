import { Chair, Sofa, FurnitureFactory } from "./interface.js";

class ModernChair extends Chair {
  sit() {
    console.log("Sitting on a modern chair");
  }
}

class ModernSofa extends Sofa {
  lieDown() {
    console.log("Lying on a modern sofa");
  }
}

export class ModernFurnitureFactory extends FurnitureFactory {
  createChair() {
    return new ModernChair();
  }
  createSofa() {
    return new ModernSofa();
  }
}
