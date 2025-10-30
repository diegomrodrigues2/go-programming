package algos

import (
	"go-code/src/heap"

	"golang.org/x/exp/constraints"
)

func MergeKLists[T constraints.Ordered](lists [][]T) ([]T, bool) {
	n := len(lists)
	heaps := make([]*heap.Heap[T], n)

	for i := 0; i < n; i++ {
		heaps[i] = heap.Heapify(lists[i])
	}

	sortedLists := make([][]T, n)

	for i := 0; i < n; i++ {
		for j := 0; j < heaps[i].Len(); j += 1 {
			item, _ := heaps[i].Pop()
			sortedLists[i] = append(sortedLists[i], item)
		}
	}

	return mergeRange(sortedLists, 0, len(sortedLists)-1), true
}

func mergeRange[T constraints.Ordered](lists [][]T, lo, hi int) []T {
	if lo == hi {
		// Podemos retornar a própria fatia; não vamos mutá-la.
		return lists[lo]
	}
	mid := lo + (hi-lo)/2
	left := mergeRange(lists, lo, mid)
	right := mergeRange(lists, mid+1, hi)
	return merge2(left, right)
}

// merge2 mescla dois slices ordenados em um novo slice ordenado (O(len(a)+len(b))).
func merge2[T constraints.Ordered](a, b []T) []T {
	out := make([]T, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	out = append(out, a[i:]...)
	out = append(out, b[j:]...)
	return out
}
