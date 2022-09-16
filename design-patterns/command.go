package main

import "fmt"

type Command interface {
	Execute()
}

type Device interface {
	On()
	Off()
}

type OnCommand struct {
	Device
}

func (c *OnCommand) Execute() {
	c.Device.On()
}

type OffCommand struct {
	Device
}

func (c *OffCommand) Execute() {
	c.Device.Off()
}

type Button struct {
	Command
}

func (b *Button) Press() {
	b.Command.Execute()
}

type TV struct {
	isRunning bool
}

func (t *TV) On() {
	if !t.isRunning {
		fmt.Println("Turning tv on")
	}
	t.isRunning = !t.isRunning
}

func (t *TV) Off() {
	if t.isRunning {
		fmt.Println("Turning tv off")
	}
	t.isRunning = !t.isRunning
}
