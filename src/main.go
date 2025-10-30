package main

import (
	"fmt"
	"go-code/src/heap"
)

func main() {
	data := []int{64, 25, 12, 22, 11}
	res, _ := heap.KLargestElements(data, 3)
	fmt.Println(res) // &{[11 12 25 22 64]}
}
