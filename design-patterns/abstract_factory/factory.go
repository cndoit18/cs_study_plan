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
	fmt.Fprintln(w.writer, "button from windows")
}

var _ Button = &linuxButton{}

type linuxButton struct {
	writer io.Writer
}

func (l *linuxButton) Click() {
	fmt.Fprintln(l.writer, "button from linux")
}

type View interface {
	Render()
}

var _ View = &windowsView{}

type windowsView struct {
	writer io.Writer
}

func (w *windowsView) Render() {
	fmt.Fprintln(w.writer, "view from windows")
}

var _ View = &linuxView{}

type linuxView struct {
	writer io.Writer
}

func (l *linuxView) Render() {
	fmt.Fprintln(l.writer, "view from linux")
}

type Factory interface {
	CreateButton() Button
	CreateView() View
}

var _ Factory = &windowsFactory{}

type windowsFactory struct {
	writer io.Writer
}

func (w *windowsFactory) CreateButton() Button {
	return &windowsButton{w.writer}
}

func (w *windowsFactory) CreateView() View {
	return &windowsView{w.writer}
}

var _ Factory = &linuxFactory{}

type linuxFactory struct {
	writer io.Writer
}

func (l *linuxFactory) CreateButton() Button {
	return &linuxButton{l.writer}
}

func (l *linuxFactory) CreateView() View {
	return &linuxView{l.writer}
}

func NewFactory(out io.Writer, os string) Factory {
	switch os {
	case "linux":
		return &linuxFactory{out}
	case "windows":
		return &windowsFactory{out}
	default:
		panic("Unsupported platform")
	}
}
