package strategy

import "golang.org/x/exp/constraints"

type Strategy[T constraints.Integer] interface {
	Execute(a, b T) T
}

type ConcreteStrategyAdd[T constraints.Integer] struct{}

func (c *ConcreteStrategyAdd[T]) Execute(a, b T) T {
	return a + b
}

type ConcreteStrategySubtract[T constraints.Integer] struct{}

func (c *ConcreteStrategySubtract[T]) Execute(a, b T) T {
	return a - b
}

type Context[T constraints.Integer] struct {
	strategy Strategy[T]
}

func (c *Context[T]) SetStrategy(s Strategy[T]) {
	c.strategy = s
}

func (c *Context[T]) Execute(a, b T) T {
	return c.strategy.Execute(a, b)
}
