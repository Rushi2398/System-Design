import { ShapeFactory } from "./factory.js";

const shape1 = ShapeFactory.createShape("Circle");
const shape2 = ShapeFactory.createShape("Rectangle");
const shape3 = ShapeFactory.createShape("Square");

shape1?.draw();
shape2?.draw();
shape3?.draw();
