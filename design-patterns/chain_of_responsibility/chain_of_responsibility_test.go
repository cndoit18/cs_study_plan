package chainofresponsibility_test

import (
	"fmt"

	chainofresponsibility "github.com/cndoit18/cs_study_plan/design-patterns/chain_of_responsibility"
)

func ExampleChain() {
	var chain chainofresponsibility.Handler
	chain = chainofresponsibility.CompoentFunc(func(a any) bool {
		fmt.Println("Begin")
		return true
	})
	chain = chain.SetNext(chainofresponsibility.CompoentFunc(func(a any) bool {
		fmt.Println("Next")
		return true
	}))
	chain = chain.SetNext(chainofresponsibility.CompoentFunc(func(a any) bool {
		fmt.Println("Failed")
		return false
	}))

	ok := chain.Handles(nil)
	fmt.Println(ok)
	// Output:
	// Begin
	// Next
	// Failed
	// false
}
