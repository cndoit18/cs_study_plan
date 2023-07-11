package memento_test

import (
	"fmt"

	"github.com/cndoit18/cs_study_plan/design-patterns/memento"
)

func ExampleMemento() {
	o := memento.Caretakers{
		Name:    "1",
		History: []memento.Memento{},
	}

	fmt.Println(o.Name)
	o.ChangeName("3")
	o.ChangeName("5")
	o.ChangeName("3")
	fmt.Println(o.Name)
	o.RollBack()
	fmt.Println(o.Name)
	o.RollBack()
	fmt.Println(o.Name)
	// Output:
	// 1
	// 3
	// 5
	// 3
}
