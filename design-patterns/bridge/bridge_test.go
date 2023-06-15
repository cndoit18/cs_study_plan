package bridge_test

import (
	"fmt"

	"github.com/cndoit18/cs_study_plan/design-patterns/bridge"
)

func ExampleShape() {
	handle := func(o bridge.Draw) {
		fmt.Println(o.Draw())
	}

	handle(&bridge.Square{
		Color: &bridge.Red{},
	})

	handle(&bridge.Square{
		Color: &bridge.Blue{},
	})

	handle(&bridge.Cricle{
		Color: &bridge.Blue{},
	})
	// Output:
	// This is a square with red color
	// This is a square with blue color
	// This is a cricle with blue color
}
