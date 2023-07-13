package visitor

import (
	"fmt"
	"io"
)

type Visitor interface {
	City(*City)
	Village(*Village)
}

type Actual struct {
	io.Writer
}

func (a *Actual) City(c *City) {
	fmt.Fprintf(a.Writer, "[City]: %s\n", c.Name)
}
func (a *Actual) Village(v *Village) {
	fmt.Fprintf(a.Writer, "[Village]: %s\n", v.Name)
}

type City struct {
	Name string
}

func (c *City) Accept(v Visitor) {
	v.City(c)
}

type Village struct {
	Name string
}

func (ve *Village) Accept(v Visitor) {
	v.Village(ve)
}
