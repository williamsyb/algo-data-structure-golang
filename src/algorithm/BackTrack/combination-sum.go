package main

import (
	"fmt"
	"sort"
)

//LeetCode
//39. 组合总和
//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
//
//说明：
//
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。
//示例 1：
//
//输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//[7],
//[2,2,3]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combination-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func combinationSum(candidates []int, target int) [][]int {
	var (
		results [][]int
		route   []int
		size    int
	)
	size = len(candidates)
	sort.Ints(candidates)
	backtrack(&route, 0, 0, target, candidates, size, &results)
	return results
}

func backtrack(route *[]int, start int, totalSum int, target int, candidates []int, size int, results *[][]int) {

	if totalSum == target {
		res := make([]int, len(*route))
		copy(res, *route)
		*results = append(*results, res)
		return
	}
	for i := start; i < size; i++ {
		totalSum += candidates[i]
		if totalSum > target {
			totalSum -= candidates[i]
			break
		}
		*route = append(*route, candidates[i])
		backtrack(route, i, totalSum, target, candidates, size, results)
		totalSum -= candidates[i]
		*route = (*route)[:len(*route)-1]

	}
	return
}

func main() {
	var r []int
	res := combinationSum([]int{2, 3, 5}, 8)
	for _, r = range res {
		fmt.Println(r)
	}
}
