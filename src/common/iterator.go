package common

type HasData interface {
	View() []any
}

type Iterator[T HasData] struct {
	src T
	i   int
}

func NewIterator[T HasData](src T) *Iterator[T] {
	return &Iterator[T]{src: src}
}

func (it *Iterator[HasData]) Next() (any, bool) {
	data := it.src.View()

	if it.i >= len(data) {
		return nil, false
	}

	v := data[it.i]
	it.i++

	return v, true
}
