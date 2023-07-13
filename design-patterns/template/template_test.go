package template_test

import (
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/template"
)

func ExampleTemplate() {
	var t template.Template = &template.Actual{template.Abstract{Writer: os.Stdout}}
	t.Print(t.Name())
	// Output:
	// [Actual], Hello,Wrold
}
