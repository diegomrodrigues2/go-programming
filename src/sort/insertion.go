package sort

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](values []T) []T {
	N := len(values)

	for i := 0; i < N; i++ {
		for j := i; j > 0 && values[j] < values[j-1]; j-- {
			values[j], values[j-1] = values[j-1], values[j]
		}
	}

	return values
}
