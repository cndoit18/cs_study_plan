package factory

import "os"

func ExampleButtonFactory() {
	button := ButtonFactory(os.Stdout, "linux")
	button.Click()

	button = ButtonFactory(os.Stdout, "windows")
	button.Click()

	// Output:
	// from linux
	// from windows
}
