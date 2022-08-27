package main

func ExampleShape() {
	var shape Shape = &Circle{}

	// 每个抽象只关心自己的实现
	shape.SetColor(&Red{})
	shape.Draw()

	shape.SetColor(&Blue{})
	shape.Draw()

	shape = &Square{}
	shape.SetColor(&Red{})
	shape.Draw()
	shape.SetColor(&Blue{})
	shape.Draw()
	// Output:
	// This is a red circle
	// This is a blue circle
	// This is a red square
	// This is a blue square
}
