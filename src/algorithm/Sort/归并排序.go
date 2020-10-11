package main

import "fmt"

const MAX_INT = int(^uint(0) >> 1)

func MergeSort(arr []int) []int {
	var (
		size  int
		mid   int
		left  []int
		right []int
	)
	size = len(arr)
	if size <= 1 {
		return arr
	}

	mid = size >> 1
	left = arr[:mid]
	right = arr[mid:]

	left = MergeSort(left[:(len(left)):(len(left))])
	right = MergeSort(right[:(len(right)):(len(right))])
	return mergeArr(left, right)
}

func mergeArr(arr01 []int, arr02 []int) (res []int) {

	var (
		i  int
		j  int
		a1 int
		a2 int
	)
	arr01 = append(arr01, MAX_INT)
	arr02 = append(arr02, MAX_INT)
	i = 0
	j = 0
	for {
		if i < len(arr01)-1 || j < len(arr02)-1 {
			a1 = arr01[i]
			a2 = arr02[j]
			if a1 < a2 {
				res = append(res, a1)
				i++
			} else {
				res = append(res, a2)
				j++
			}
		} else {
			break
		}
	}
	return
}

func change(arr []int) []int {
	arr = append(arr, 100)
	return arr
}

func main() {

	var arr = []int{4, 6, 9, 2, 1, 3, 11, 32}
	//fmt.Println(arr)
	//fmt.Println(change(arr[:3]))
	//fmt.Println(arr)

	//fmt.Println(mergeArr([]int{-1}, []int{1}))
	fmt.Println(MergeSort(arr))
}
