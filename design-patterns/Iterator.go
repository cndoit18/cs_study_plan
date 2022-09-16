package main

type Iterator[T any] interface {
	Next() bool
	Value() T
}

type IterSlice[T any] struct {
	i    int
	data []T
}

func (iter *IterSlice[T]) Next() bool {
	return iter.i < len(iter.data)
}

func (iter *IterSlice[T]) Value() T {
	defer func() { iter.i += 1 }()
	return iter.data[iter.i]
}
