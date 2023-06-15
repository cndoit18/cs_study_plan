package factory_test

import (
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/factory"
)

func ExampleButtonFactory() {
	button := factory.ButtonFactory(os.Stdout, "linux")
	button.Click()

	button = factory.ButtonFactory(os.Stdout, "windows")
	button.Click()

	// Output:
	// from linux
	// from windows
}
