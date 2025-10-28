package sort

import "golang.org/x/exp/constraints"

// MergeSort sorts values in ascending order (stable) and returns the same slice.
func MergeSort[T constraints.Ordered](a []T) []T {
	if len(a) < 2 {
		return a
	}
	aux := make([]T, len(a)) // one-time aux buffer
	msort(a, aux, 0, len(a)-1)
	return a
}

func msort[T constraints.Ordered](a, aux []T, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	msort(a, aux, lo, mid)
	msort(a, aux, mid+1, hi)

	// Small optimization: already in order
	if a[mid] <= a[mid+1] {
		return
	}
	merge(a, aux, lo, mid, hi)
}

func merge[T constraints.Ordered](a, aux []T, lo, mid, hi int) {
	// Copy a[lo:hi+1] into aux
	copy(aux[lo:hi+1], a[lo:hi+1])

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		switch {
		case i > mid:
			a[k] = aux[j]
			j++
		case j > hi:
			a[k] = aux[i]
			i++
		case aux[j] < aux[i]:
			a[k] = aux[j]
			j++
		default:
			a[k] = aux[i]
			i++
		}
	}
}

func MergeSortIter[T constraints.Ordered](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	aux := make([]T, n)
	for sz := 1; sz < n; sz += sz {
		for lo := 0; lo < n-sz; lo += 2 * sz {
			mid := lo + sz - 1
			hi := lo + 2*sz - 1
			if hi >= n {
				hi = n - 1
			}
			if a[mid] <= a[mid+1] { // skip if already ordered
				continue
			}
			merge(a, aux, lo, mid, hi)
		}
	}
	return a
}
