import { SimpleShapeFactory } from "./simpleShapeFactory.js";

const factory = new SimpleShapeFactory();
const shape1 = factory.createShape("Circle");
const shape2 = factory.createShape("Rectangle");
const shape3 = factory.createShape("Square");

shape1?.draw();
shape2?.draw();
shape3?.draw();
