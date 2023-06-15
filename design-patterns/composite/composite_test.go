package composite_test

import (
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/composite"
)

func ExampleComposite() {
	var container composite.Component
	container = &composite.Columns{}

	rows := &composite.Rows{}
	rows.Add(&composite.Dot{}, &composite.Dot{})

	container.Add(rows, &composite.Dot{}, &composite.Dot{}, rows)
	container.Draw(os.Stdout)
	// Output:
	// ⭐️⭐️
	// ⭐️
	// ⭐️
	// ⭐️⭐️
}
