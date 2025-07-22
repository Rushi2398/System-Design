package main

type SimpleShapeFactory struct{}

func (s SimpleShapeFactory) CreateShape(shape string) shapeInterface {
	switch shape {
	case "Circle":
		return Circle{}
	case "Rectangle":
		return Rectangle{}
	case "Square":
		return Square{}
	default:
		return nil
	}
}
