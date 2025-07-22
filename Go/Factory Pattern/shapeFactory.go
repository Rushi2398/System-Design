package main

type ShapeFactory interface {
	CreateShape(shape string) shapeInterface
}
