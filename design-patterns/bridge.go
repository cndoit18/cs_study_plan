package main

import "fmt"

type Shape interface {
	SetColor(Color)
	Draw()
}

type Color interface {
	String() string
}

type Circle struct {
	Color
}

func (c *Circle) SetColor(color Color) {
	c.Color = color
}

func (c *Circle) Draw() {
	fmt.Printf("This is a %s circle\n", c.Color.String())
}

type Square struct {
	Color
}

func (s *Square) SetColor(color Color) {
	s.Color = color
}

func (s *Square) Draw() {
	fmt.Printf("This is a %s square\n", s.Color.String())
}

type Red struct{}

func (r *Red) String() string {
	return "red"
}

type Blue struct{}

func (b *Blue) String() string {
	return "blue"
}
