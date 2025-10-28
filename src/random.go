package main

import (
	"math/rand"
)

func Discrete(a []float64) int {
	r := rand.Float64()

	sum := 0.0

	for i, p := range a {
		sum += p
		if sum >= r {
			return i
		}
	}

	return -1
}

func Shuffle(a []float64) {
	n := len(a)

	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		a[i], a[r] = a[r], a[i]
	}
}
