package main

import (
	"fmt"
)

func PartitionWithFirstElementAsPivot(arr []int, lo, hi int) int {
	pivot := arr[lo]
	for j := hi; j > lo; j-- {
		if arr[j] > pivot {
			arr[j], arr[hi] = arr[hi], arr[j]
			hi--
		}

	}
	arr[hi], arr[lo] = arr[lo], arr[hi]

	return lo

}

func PartitionAslastElementAsPivot(arr []int, lo, hi int) int {
	pivot := arr[hi]
	for j := lo; j < hi; j++ {
		if arr[j] < pivot {
			arr[j], arr[lo] = arr[lo], arr[j]
			lo++
		}

	}
	arr[hi], arr[lo] = arr[lo], arr[hi]

	return lo

}

func QuickSort(arr []int, l, hi int) {
	if l > hi {
		return
	}
	p := PartitionWithFirstElementAsPivot(arr, l, hi)
	//p := PartitionAslastElementAsPivot(arr, l, hi)
	QuickSort(arr, l, p-1)
	QuickSort(arr, p+1, hi)

}

func main() {
	fmt.Println("Hello, playground")
	arr := []int{5, 3, 2, 7, 6, 8}
	var l = 0
	var hi = len(arr) - 1

	QuickSort(arr, l, hi)
	fmt.Println(arr)
}
