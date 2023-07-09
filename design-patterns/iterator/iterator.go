package iterator

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type List[T any] struct {
	data    []T
	current int
}

func (l *List[T]) HasNext() bool {
	return l.current < len(l.data)
}

func (l *List[T]) Next() T {
	cur := l.current
	l.current++
	return l.data[cur]
}

func NewIterator[T any](array ...T) Iterator[T] {
	return &List[T]{
		data:    array,
		current: 0,
	}
}
