package command_test

import (
	"bytes"
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/command"
)

func ExampleCommand() {
	buffer := &bytes.Buffer{}
	commands := []command.Interface{}
	commands = append(commands,
		&command.EchoCommand{"Hello", buffer})
	commands = append(commands,
		&command.PrintCommand{os.Stdout, buffer})

	for _, c := range commands {
		c.Execute()
	}
	// Output:
	// Hello
}
