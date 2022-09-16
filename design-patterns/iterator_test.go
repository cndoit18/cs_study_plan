package main

import "fmt"

func ExampleIterator() {
	iter := &IterSlice[int]{
		data: []int{1, 2, 3, 4, 5},
	}
	for iter.Next() {
		fmt.Println(iter.Value())
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
