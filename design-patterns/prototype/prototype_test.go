package prototype_test

import (
	"fmt"
	"reflect"

	"github.com/cndoit18/cs_study_plan/design-patterns/prototype"
)

func ExamplePrototype() {
	p := prototype.People{
		Name:   "cndoit18",
		Age:    28,
		Gender: prototype.Male,
	}

	fmt.Println(p.Clone() == &p)
	fmt.Println(reflect.DeepEqual(p.Clone(), &p))
	// Output:
	// false
	// true
}
