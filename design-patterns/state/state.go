package state

import (
	"fmt"
	"io"
)

type State interface {
	Open()
	Close()
}

func NewCompse(out io.Writer) *Compose {
	state := &CloseState{
		Writer: out,
	}
	c := &Compose{state: state}
	state.C = c
	return c
}

type Compose struct {
	state State
}

func (c *Compose) ChangeState(s State) {
	c.state = s
}

func (c *Compose) Open() {
	c.state.Open()
}

func (c *Compose) Close() {
	c.state.Close()
}

type OpenState struct {
	C *Compose
	io.Writer
}

func (o *OpenState) Open() {
	fmt.Fprintln(o.Writer, "Already open, no need to open again")
}

func (o *OpenState) Close() {
	fmt.Fprintln(o.Writer, "Closing")
	o.C.ChangeState(&CloseState{C: o.C, Writer: o.Writer})
}

type CloseState struct {
	C *Compose
	io.Writer
}

func (o *CloseState) Open() {
	fmt.Fprintln(o.Writer, "Opening")
	o.C.ChangeState(&OpenState{C: o.C, Writer: o.Writer})
}

func (o *CloseState) Close() {
	fmt.Fprintln(o.Writer, "Already closed, no need to close again")
}
