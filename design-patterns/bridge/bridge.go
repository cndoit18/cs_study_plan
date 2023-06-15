package bridge

import "fmt"

type Draw interface {
	Draw() string
}

type Shape interface {
	Shape() string
}

var (
	_ Shape = &Cricle{}
	_ Shape = &Square{}
	_ Color = &Red{}
	_ Color = &Blue{}
)

type Cricle struct {
	Color
}

func (c *Cricle) Shape() string {
	return "cricle"
}

func (c *Cricle) Draw() string {
	return fmt.Sprintf("This is a %s with %s color", c.Shape(), c.Color.Color())
}

type Square struct {
	Color
}

func (s *Square) Shape() string {
	return "square"
}

func (s *Square) Draw() string {
	return fmt.Sprintf("This is a %s with %s color", s.Shape(), s.Color.Color())
}

type Color interface {
	Color() string
}

type Red struct {
}

func (r *Red) Color() string {
	return "red"
}

type Blue struct {
}

func (b *Blue) Color() string {
	return "blue"
}
