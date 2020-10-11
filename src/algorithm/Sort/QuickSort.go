package main

import "fmt"

func partition2(arr []int, low int, high int) int {
	var (
		pivot int
	)
	pivot = arr[low]
	for low < high {
		for low < high && pivot < arr[high] {
			high--
		}
		arr[low] = arr[high]
		for low < high && pivot >= arr[low] {
			low++
		}
		arr[high] = arr[low]

	}
	arr[low] = pivot
	return low
}

func QuickSort2(arr []int, start int, end int) []int {
	var (
		pivotLoc int
	)
	if start < end {
		pivotLoc = partition2(arr, start, end)
		QuickSort2(arr, start, pivotLoc)
		QuickSort2(arr, pivotLoc+1, end)
	}
	return arr
}

func main() {
	list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	QuickSort2(list, 0, len(list)-1)
	fmt.Print(list)
}
