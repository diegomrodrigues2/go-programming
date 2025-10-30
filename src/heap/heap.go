package heap

import (
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	data []T
}

// Heapify builds a heap in-place from a (uses the backing array of a).
// If you don't want to mutate the input slice, copy it first.
func Heapify[T constraints.Ordered](a []T) *Heap[T] {
	h := &Heap[T]{data: a}
	// start from last internal node
	for i := (len(h.data) - 2) / 2; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *Heap[T]) Len() int      { return len(h.data) }
func (h *Heap[T]) IsEmpty() bool { return len(h.data) == 0 }

func (h *Heap[T]) Insert(x T) {
	h.data = append(h.data, x)
	h.siftUp(len(h.data) - 1)
}

// Peek returns the minimum element without removing it.
func (h *Heap[T]) Peek() (T, bool) {
	var zero T
	if len(h.data) == 0 {
		return zero, false
	}
	return h.data[0], true
}

// Pop removes and returns the minimum element.
func (h *Heap[T]) Pop() (T, bool) {
	var zero T
	n := len(h.data)
	if n == 0 {
		return zero, false
	}
	min := h.data[0]
	// move last to root and shrink
	last := h.data[n-1]
	h.data[n-1] = zero // help GC
	h.data = h.data[:n-1]
	if len(h.data) > 0 {
		h.data[0] = last
		h.siftDown(0)
	}
	return min, true
}

// ----- internals (min-heap) -----

func (h *Heap[T]) siftUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		// min-heap: bubble up while child < parent
		if h.data[i] < h.data[p] {
			h.data[i], h.data[p] = h.data[p], h.data[i]
			i = p
		} else {
			break
		}
	}
}

func (h *Heap[T]) siftDown(i int) {
	n := len(h.data)
	for {
		l := 2*i + 1 // left child (0-based)
		r := 2*i + 2 // right child
		smallest := i

		if l < n && h.data[l] < h.data[smallest] {
			smallest = l
		}
		if r < n && h.data[r] < h.data[smallest] {
			smallest = r
		}
		if smallest == i {
			return
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

// NLargest returns the n largest elements from a, in descending order.
// Runs in O(len(a) log n). Does not mutate the input slice a.
func NLargest[T constraints.Ordered](a []T, n int) []T {
	if n <= 0 || len(a) == 0 {
		return nil
	}
	if n > len(a) {
		n = len(a)
	}

	// 1) Build a min-heap from the first n items (copy to avoid mutating a[:n]).
	buf := make([]T, n)
	copy(buf, a[:n])
	h := &Heap[T]{data: buf}
	for i := (n - 2) / 2; i >= 0; i-- { // last internal node
		h.siftDown(i)
	}

	// 2) Maintain top-n against the remaining items.
	for i := n; i < len(a); i++ {
		if a[i] > h.data[0] { // compare with heap min
			h.data[0] = a[i]
			h.siftDown(0)
		}
	}

	// 3) Extract to get descending order (largest -> smallest).
	res := make([]T, n)
	for i := n - 1; i >= 0; i-- {
		res[i], _ = h.Pop() // Pop() is O(log n) on our min-heap
	}
	return res
}

func HeapSort[T constraints.Ordered](a []T) []T {
	n := len(a)
	h := Heapify(a)

	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i], _ = h.Pop() // Pop() is O(log n) on our min-heap
	}
	return res
}

func NextGreatElement[T constraints.Ordered](a []T, elem T) (T, bool) {
	n := len(a)
	h := Heapify(a)

	for i := 0; i < n; i++ {
		item, _ := h.Pop()
		if item > elem {
			return item, true
		}
	}

	var out T
	return out, false
}

func NextGreatElements[T constraints.Ordered](a []T, elem T) ([]T, bool) {
	n := len(a)
	h := Heapify(a)

	res := make([]T, 0)

	for i := 0; i < n; i++ {
		item, _ := h.Pop()
		if item > elem {
			res = append(res, item)
		}
	}

	return res, true
}

func KLargestElements[T constraints.Ordered](a []T, k int) ([]T, bool) {

	aux := make([]T, k)
	for i := 0; i < k; i++ {
		aux[i] = a[i]
	}

	heap := Heapify(aux)

	for i := k; i < len(a); i++ {
		min, _ := heap.Peek()
		if a[i] > min {
			heap.Pop()
			heap.Insert(a[i])
		}
	}

	res := make([]T, k)
	for i := 0; i < k; i++ {
		res[i], _ = heap.Pop() // Pop() is O(log n) on our min-heap
	}

	return res, true
}
