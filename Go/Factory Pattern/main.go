package main

func main() {
	factory := SimpleShapeFactory{}

	shape1 := factory.CreateShape("Circle")
	shape2 := factory.CreateShape("Rectangle")
	shape3 := factory.CreateShape("Square")

	if shape1 != nil {
		shape1.Draw()
	}
	if shape2 != nil {
		shape2.Draw()
	}
	if shape3 != nil {
		shape3.Draw()
	}
}
