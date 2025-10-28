package main

import (
	"fmt"
	"go-code/src/sort"
)

func main() {
	data := []int{64, 25, 12, 22, 11}
	sort.QuickSort(data)
	fmt.Println(data) // [11 12 22 25 64]
}
