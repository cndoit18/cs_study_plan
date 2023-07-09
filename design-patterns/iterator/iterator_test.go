package iterator_test

import (
	"fmt"

	"github.com/cndoit18/cs_study_plan/design-patterns/iterator"
)

func ExampleIterator() {
	iter := iterator.NewIterator[int](1, 2, 3, 4, 5)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
