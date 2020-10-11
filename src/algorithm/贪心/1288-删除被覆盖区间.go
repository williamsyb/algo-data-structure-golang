package main

import (
	"fmt"
	"sort"
)

type array [][]int

func (a array) Len() int { return len(a) }
func (a array) Less(i, j int) bool {
	if a[i][0] < a[j][0] {
		return true
	} else if a[i][0] > a[j][0] {
		return false
	} else {
		if a[i][1] < a[j][1] {
			return false
		} else {
			return true
		}
	}
}
func (a array) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func removeCoveredIntervals(intervals [][]int) int {
	var (
		arr   array
		left  int
		right int
		item  []int
		i     int
		res   int
	)
	arr = array(intervals)
	sort.Sort(arr)
	//拿到头一个元素的 起始值
	left = arr[0][0]
	right = arr[0][1]
	res = 0
	for i = 1; i < len(arr); i++ {
		item = arr[i]
		if item[0] <= right && item[1] <= right && item[0] >= left {
			res++
		} else if item[0] <= right && item[1] >= right {
			//left = left
			right = item[1]
		} else {
			left = item[0]
			right = item[1]
		}

	}
	return len(arr) - res
}

func main() {
	arr := array{{1, 4}, {3, 6}, {2, 8}, {12, 4}, {12, 3}, {12, 5}, {10, 1}, {122, 3}, {100, 12}, {42, 5}}

	sort.Sort(arr)
	fmt.Println(arr)
	fmt.Printf("原始长度：%v， 剩余长度：%v", len(arr), removeCoveredIntervals(arr))
}
