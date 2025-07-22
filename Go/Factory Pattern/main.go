package main

func main() {
	shape1 := Shape("Circle")
	shape2 := Shape("Rectangle")
	shape3 := Shape("Square")

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
