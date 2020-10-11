package main

import "fmt"

func permutation(arr []int) [][]int {
	var (
		results *[][]int
		route   *[]int
		size    int
		flags   *[]bool
		i       int
	)

	results = new([][]int)
	*results = [][]int{}
	route = new([]int)
	*route = []int{}
	flags = new([]bool)
	size = len(arr)
	*flags = make([]bool, size)

	for i = 0; i < size; i++ {
		(*flags)[i] = false
	}
	backtrackPer(arr, results, route, size, flags)
	return *results
}

func backtrackPer(arr []int, results *[][]int, route *[]int, size int, flags *[]bool) {
	var res []int
	var i int
	if len(*route) == size {
		res = make([]int, len(*route))
		copy(res, *route)
		*results = append(*results, res)
		return
	}
	for i = 0; i < size; i++ {
		if (*flags)[i] == false {
			*route = append(*route, arr[i])
			(*flags)[i] = true
			backtrackPer(arr, results, route, size, flags)
			*route = (*route)[:len((*route))-1]
			(*flags)[i] = false
		}
	}
}

func main() {
	var (
		results [][]int
		val     []int
	)
	results = permutation([]int{1, 2, 3, 4})

	for _, val = range results {
		fmt.Println(val)
	}
}
