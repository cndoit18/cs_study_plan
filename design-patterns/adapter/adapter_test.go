package adapter_test

import (
	"fmt"
	"reflect"

	"github.com/cndoit18/cs_study_plan/design-patterns/adapter"
)

func ExampleNewRouadAdapater() {
	handle := func(round adapter.Round) {
		fmt.Printf("[%s] getRadius: %d\n", reflect.TypeOf(round).Elem().Name(), round.GetRadius())
	}

	rounds := []adapter.Round{
		&adapter.RoundHole{Radius: 1},
		&adapter.RoundPeg{Radius: 2},
		adapter.NewRouadAdapater(&adapter.SquarePeg{Width: 3}),
	}
	for _, v := range rounds {
		handle(v)
	}
	// Output:
	// [RoundHole] getRadius: 1
	// [RoundPeg] getRadius: 2
	// [roundAdapter] getRadius: 3
}
