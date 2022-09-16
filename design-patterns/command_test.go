package main

func ExampleCommand() {
	tv := &TV{}

	onCommand := &OnCommand{
		Device: tv,
	}

	offCommand := &OffCommand{
		Device: tv,
	}

	onButton := &Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &Button{
		Command: offCommand,
	}
	offButton.Press()
	// Output:
	// Turning tv on
	// Turning tv off
}
