package main

import "fmt"

func ExampleColorFactory() {
	factory := ColorFactory{
		colors: map[string]Color{},
	}
	// 共享同一个底层数据
	r0 := factory.GetColor("red")
	fmt.Printf("color: %q\n", r0.String())
	o := r0.(*PtrColor)
	o.color = "blue"
	fmt.Printf("color: %q\n", r0.String())
	// Output:
	// color: "red"
	// color: "blue"
}
