package factory

import (
	"fmt"
	"io"
)

type Button interface {
	Click()
}

var _ Button = &windowsButton{}

type windowsButton struct {
	writer io.Writer
}

func (w *windowsButton) Click() {
	fmt.Fprintln(w.writer, "from windows")
}

var _ Button = &linuxButton{}

type linuxButton struct {
	writer io.Writer
}

func (l *linuxButton) Click() {
	fmt.Fprintln(l.writer, "from linux")
}

func ButtonFactory(out io.Writer, os string) Button {
	switch os {
	case "linux":
		return &linuxButton{out}
	case "windows":
		return &windowsButton{out}
	default:
		panic("Unsupported platform")
	}
}
