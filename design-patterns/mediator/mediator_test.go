package mediator_test

import (
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/mediator"
)

func ExampleMediator() {
	root := &mediator.Room{
		Writer: os.Stdout,
	}

	x, y := mediator.NewUser("x", root), mediator.NewUser("y", root)
	x.SendMessage("hello")
	y.SendMessage("hello")
	// Output:
	// [x]: hello
	// [y]: hello
}
