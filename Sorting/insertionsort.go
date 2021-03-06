package main

import (
	"fmt"
)
func insertionSort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i > -1 && A[i] > key {
			A[i+1] = A[i]
			i -= 1
		}
		A[i+1] = key
	}
	return A
}
func main() {
	arr := []int {2,1,5,8,4}
	a:= insertionSort(arr)
	fmt.Println(a)
}
