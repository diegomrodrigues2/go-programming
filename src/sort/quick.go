package sort

import "golang.org/x/exp/constraints"

// QuickSort sorts the slice a in ascending order in place and returns it.
// Works for any ordered type (ints, floats, strings...).
func QuickSort[T constraints.Ordered](a []T) []T {
	if len(a) < 2 {
		return a
	}
	quickSort(a, 0, len(a)-1)
	return a
}

func quickSort[T constraints.Ordered](a []T, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

// partition rearranges the slice so that a[lo:p] <= a[p] <= a[p+1:hi]
// and returns the final index of the pivot.
func partition[T constraints.Ordered](a []T, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
