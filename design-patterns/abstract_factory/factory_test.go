package factory_test

import (
	"os"

	factory "github.com/cndoit18/cs_study_plan/design-patterns/abstract_factory"
)

func ExampleFactory() {
	f := factory.NewFactory(os.Stdout, "windows")
	f.CreateButton().Click()
	f.CreateView().Render()

	f = factory.NewFactory(os.Stdout, "linux")
	f.CreateButton().Click()
	f.CreateView().Render()
	// Output:
	// button from windows
	// view from windows
	// button from linux
	// view from linux
}
