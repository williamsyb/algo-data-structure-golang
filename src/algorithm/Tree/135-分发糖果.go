package main

import "fmt"

/*
135. 分发糖果
难度 ：hard
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

示例 1:

输入: [1,0,2]
输出: 5
解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
示例 2:

输入: [1,2,2]
输出: 4
解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
*/

func candy(ratings []int) int {
	var (
		size int
		i    int
		nums []int
		res  int
	)
	size = len(ratings)
	if size < 2 {
		return size
	}
	nums = []int{}
	for i = 0; i < size; i++ {
		nums = append(nums, 1)
	}

	for i = 1; i < size; i++ {
		if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		}
	}
	for i = size - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && nums[i-1] <= nums[i] {
			nums[i-1] = nums[i] + 1
		}
	}
	res = 0
	for _, num := range nums {
		res += num

	}
	return res
}
func main() {
	fmt.Println(candy([]int{1, 2, 2}))
}
