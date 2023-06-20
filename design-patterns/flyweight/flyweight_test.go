package flyweight_test

import (
	"fmt"

	"github.com/cndoit18/cs_study_plan/design-patterns/flyweight"
)

func ExampleFlyweight() {
	trees := flyweight.GenerateTrees(10)
	for _, v := range trees.Trees {
		fmt.Println(v.Name())
	}
	// Output:
	// tree
	// tree
	// tree
	// tree
	// tree
	// tree
	// tree
	// tree
	// tree
	// tree
}
