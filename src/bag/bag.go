package bag

import "go-code/src/common"

type Bag[T any] struct {
	items []T
}

func New[T any]() *Bag[T] {
	return &Bag[T]{}
}

func (b *Bag[T]) Items() []T {
	return b.items
}

func (b *Bag[T]) View() []any {
	out := make([]any, len(b.items))
	for i, item := range b.items {
		out[i] = item
	}
	return out
}

func (b *Bag[T]) Add(x T) {
	b.items = append(b.items, x)
}

func (b *Bag[T]) IsEmpty() bool {
	return len(b.items) == 0
}

func (b *Bag[T]) Len() int {
	return len(b.items)
}

func (b *Bag[T]) Slice() []T {
	out := make([]T, len(b.items))
	copy(out, b.items)
	return out
}

func (b *Bag[T]) Iterator() *common.Iterator[*Bag[T]] {
	return common.NewIterator(b)
}
