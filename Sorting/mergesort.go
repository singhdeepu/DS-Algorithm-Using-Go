package main

import (
	"fmt"
)

func merge(a, b []int) []int {
	resArray := make([]int, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			resArray[i+j] = a[i]
			i++

		} else {
			resArray[i+j] = b[j]

			j++

		}
	}
	for i < len(a) {
		resArray[i+j] = a[i]
		i++
	}

	for j < len(b) {
		resArray[i+j] = b[j]
		j++

	}
	
	return resArray

}
func mergeSort(items []int) []int {

	if len(items) < 2 {

		return items
	}

	middle := len(items) / 2

	a := mergeSort(items[:middle])
	b := mergeSort(items[middle:])

	return merge(a, b)

}
func main() {
	fmt.Println("Hello, playground")
	
	inputSlice := []int { 3,2,5,7,9,1}
	
	sortedSlice := mergeSort(inputSlice)
	
	fmt.Println(sortedSlice)
}
