package factory

import "os"

func ExampleFactory() {
	factory := NewFactory(os.Stdout, "windows")
	factory.CreateButton().Click()
	factory.CreateView().Render()

	factory = NewFactory(os.Stdout, "linux")
	factory.CreateButton().Click()
	factory.CreateView().Render()
	// Output:
	// button from windows
	// view from windows
	// button from linux
	// view from linux
}
