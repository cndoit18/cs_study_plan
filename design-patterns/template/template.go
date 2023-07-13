package template

import (
	"fmt"
	"io"
)

type Template interface {
	Print(string)
	Name() string
}

type Abstract struct {
	io.Writer
}

func (a *Abstract) Print(name string) {
	fmt.Fprintf(a.Writer, "[%s], Hello,Wrold\n", name)
}

func (a *Abstract) Name() string {
	panic("not implemented") // TODO: Implement
}

type Actual struct {
	Abstract
}

func (a *Actual) Name() string {
	return "Actual"
}
