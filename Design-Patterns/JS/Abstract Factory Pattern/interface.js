// Abstract Product Interfaces
export class Chair {
  sit() {
    throw new Error("Abstract method!");
  }
}

export class Sofa {
  lieDown() {
    throw new Error("Abstract method!");
  }
}

// Abstract Factory Interface
export class FurnitureFactory {
  createChair() {
    throw new Error("Not implemented");
  }
  createSofa() {
    throw new Error("Not implemented");
  }
}
