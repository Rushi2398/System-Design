import { ShapeFactory } from "./factory.js";
import { Circle } from "./circle.js";
import { Rectangle } from "./rectangle.js";
import { Square } from "./square.js";

export class SimpleShapeFactory extends ShapeFactory {
  createShape(type) {
    switch (type) {
      case "Circle":
        return new Circle();
      case "Rectangle":
        return new Rectangle();
      case "Square":
        return new Square();
      default:
        return null;
    }
  }
}
