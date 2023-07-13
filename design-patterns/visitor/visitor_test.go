package visitor_test

import (
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/visitor"
)

func ExampleVisitor() {
	data := []interface{ Accept(visitor.Visitor) }{
		&visitor.City{Name: "abc"}, &visitor.Village{Name: "qq"},
	}
	v := &visitor.Actual{Writer: os.Stdout}

	for _, d := range data {
		d.Accept(v)
	}
	// Output:
	// [City]: abc
	// [Village]: qq
}
