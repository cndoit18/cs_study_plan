package state_test

import (
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/state"
)

func ExampleState() {
	c := state.NewCompse(os.Stdout)
	c.Close()

	c.Open()
	c.Open()
	c.Close()
	// Output:
	// Already closed, no need to close again
	// Opening
	// Already open, no need to open again
	// Closing
}
