package singleton_test

import (
	"fmt"

	"github.com/cndoit18/cs_study_plan/design-patterns/singleton"
)

func ExampleGetInstanceFunc() {
	getinstance := singleton.GetInstanceFunc()
	instance := getinstance("cndoit18")
	fmt.Println(instance.Value)
	instance = getinstance("cndoit17")
	fmt.Println(instance.Value)
	// Output:
	// cndoit18
	// cndoit18
}
