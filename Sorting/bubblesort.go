package main

import "fmt"

func bubbleSort(A []int) []int {
	n := len(A)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
			fmt.Println("j:", j)
		}
		fmt.Println("i:", i)
	}
	return A
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	A = bubbleSort(A)
	fmt.Println("\n After Bubble Sorting")

	fmt.Println(A)
}
