package stack

// FixedSizeStack é uma pilha com capacidade máxima fixa.
type FixedSizeStack[T any] struct {
	items    []T
	capacity int
}

// NewFixedSizeStack cria uma pilha com a capacidade indicada.
func NewFixedSizeStack[T any](capacity int) *FixedSizeStack[T] {
	return &FixedSizeStack[T]{
		items:    make([]T, 0, capacity),
		capacity: capacity,
	}
}

// IsEmpty informa se a pilha está vazia.
func (s *FixedSizeStack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size retorna o número de elementos na pilha.
func (s *FixedSizeStack[T]) Size() int {
	return len(s.items)
}

// Push adiciona um elemento no topo da pilha.
// Retorna false se a pilha estiver cheia.
func (s *FixedSizeStack[T]) Push(v T) bool {
	if len(s.items) >= s.capacity {
		return false // pilha cheia
	}
	s.items = append(s.items, v)
	return true
}

// Pop remove e retorna o elemento mais recente.
// Retorna (zero, false) se a pilha estiver vazia.
func (s *FixedSizeStack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items[last] = zero // evita vazamento de referência
	s.items = s.items[:last]
	return v, true
}

// Peek retorna o elemento no topo da pilha sem removê-lo.
func (s *FixedSizeStack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}
