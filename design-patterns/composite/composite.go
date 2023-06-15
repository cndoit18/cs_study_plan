package composite

import (
	"fmt"
	"io"
)

type Component interface {
	Draw(io.Writer)
	Add(...Component)
}

type Dot struct {
}

func (d *Dot) Draw(out io.Writer) {
	fmt.Fprint(out, "⭐️")
}

func (d *Dot) Add(c ...Component) {
	panic("dot has no way to add other components")
}

type Rows struct {
	components []Component
}

func (r *Rows) Draw(out io.Writer) {
	for _, c := range r.components {
		c.Draw(out)
	}
}

func (r *Rows) Add(c ...Component) {
	r.components = append(r.components, c...)
}

type Columns struct {
	components []Component
}

func (c *Columns) Draw(out io.Writer) {
	for _, c := range c.components {
		c.Draw(out)
		fmt.Fprintln(out)
	}
}

func (cs *Columns) Add(c ...Component) {
	cs.components = append(cs.components, c...)
}
