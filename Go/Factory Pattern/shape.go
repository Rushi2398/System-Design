package main

func Shape(shape string) shapeInterface {
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
