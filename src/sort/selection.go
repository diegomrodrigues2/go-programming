package sort

import "golang.org/x/exp/constraints"

// SelectionSort sorts a slice of ordered values in ascending order in place
// and returns the same slice for convenience.
func SelectionSort[T constraints.Ordered](values []T) []T {
	n := len(values)

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if values[j] < values[min] {
				min = j
			}
		}
		// swap values[i] and values[min]
		values[i], values[min] = values[min], values[i]
	}
	return values
}
