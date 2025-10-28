package linkedlist

type Node[T any] struct {
	item T
	next *Node[T]
}

func New[T any](v T) *Node[T] {
	return &Node[T]{item: v}
}

// Build a Linked List from a array of type T items
func Build[T any](values []T) *Node[T] {
	if len(values) == 0 {
		return nil
	}

	head := &Node[T]{item: values[0]}
	tail := head

	for i := 1; i < len(values); i++ {
		n := &Node[T]{item: values[i]}
		tail.next = n
		tail = n
	}

	return head
}

// Add to the begining
func (l *Node[T]) Prepend(v T) (head *Node[T], ok bool) {
	if l == nil {
		return &Node[T]{item: v}, true
	}

	newHead := &Node[T]{item: v, next: l}
	return newHead, true
}

func (l *Node[T]) Insert(v T) (head *Node[T], ok bool) {
	if l == nil {
		return &Node[T]{item: v}, true
	}

	head = l
	current := l

	for current.next != nil {
		current = current.next
	}

	current.next = &Node[T]{item: v}

	return head, true
}

// Remove from the beggining
func (l *Node[T]) Pull() (head *Node[T], ok bool) {
	if l == nil {
		return nil, false
	}

	return l.next, true
}

// Iterator — returns a channel of items for easy traversal
func (l *Node[T]) Iterator() <-chan T {
	ch := make(chan T)

	go func() {
		for n := l; n != nil; n = n.next {
			ch <- n.item
		}
		close(ch)
	}()

	return ch
}

type Iter[T any] struct {
	n *Node[T]
}

func (l *Node[T]) Iter() Iter[T] {
	return Iter[T]{n: l}
}

func (it *Iter[T]) Next() (v T, ok bool) {
	if it.n == nil {
		var zero T
		return zero, false
	}
	v = it.n.item
	it.n = it.n.next
	return v, true
}

// InsertAt insere v na posição pos (0-based).
func (l *Node[T]) InsertAt(v T, pos int) (head *Node[T], ok bool) {
	// índices negativos não são válidos
	if pos < 0 {
		return l, false
	}

	// lista vazia: só é válido inserir em pos==0
	if l == nil {
		if pos == 0 {
			return &Node[T]{item: v}, true
		}
		return nil, false
	}

	// inserir no começo (pos 0)
	if pos == 0 {
		return &Node[T]{item: v, next: l}, true
	}

	head = l
	prev := (*Node[T])(nil)
	cur := l
	i := 0

	// avançar até a posição desejada (ou fim)
	for cur != nil && i < pos {
		prev = cur
		cur = cur.next
		i++
	}

	// se não alcançou pos, pos > len(lista) → falha
	if i != pos || prev == nil {
		return head, false
	}

	// insere entre prev e cur (cur pode ser nil: insere no fim)
	node := &Node[T]{item: v, next: cur}
	prev.next = node
	return head, true
}
