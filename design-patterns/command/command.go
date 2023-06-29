package command

import (
	"fmt"
	"io"
)

type Interface interface {
	Execute()
}

type PrintCommand struct {
	Out io.Writer
	In  io.Reader
}

func (p PrintCommand) Execute() {
	io.Copy(p.Out, p.In)
}

type EchoCommand struct {
	Message string
	Out     io.Writer
}

func (e EchoCommand) Execute() {
	fmt.Fprint(e.Out, e.Message)
}
