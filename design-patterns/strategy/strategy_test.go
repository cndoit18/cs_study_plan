package strategy_test

import (
	"fmt"

	"github.com/cndoit18/cs_study_plan/design-patterns/strategy"
)

func ExampleStrategy() {
	c := strategy.Context[int]{}
	c.SetStrategy(&strategy.ConcreteStrategyAdd[int]{})
	fmt.Println(c.Execute(1, 2))

	c.SetStrategy(&strategy.ConcreteStrategySubtract[int]{})
	fmt.Println(c.Execute(1, 2))
	// Output:
	// 3
	// -1
}
