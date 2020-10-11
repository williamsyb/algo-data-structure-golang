package main

//40. 组合总和 II
//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的每个数字在每个组合中只能使用一次。
//
//说明：
//
//所有数字（包括目标数）都是正整数。
//解集不能包含重复的组合。
//示例 1:
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//[1, 7],
//[1, 2, 5],
//[2, 6],
//[1, 1, 6]
//]
//示例 2:
//
//输入: candidates = [2,5,2,1,2], target = 5,
//所求解集为:
//[
//  [1,2,2],
//  [5]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combination-sum-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var (
		results *[][]int
		route   *[]int
	)
	sort.Ints(candidates)
	fmt.Println(candidates)
	results = new([][]int)
	*results = [][]int{}
	route = new([]int)
	*route = []int{}
	backTrack(route, results, len(candidates), 0, candidates, target, 0)
	return *results
}

func backTrack(route *[]int, results *[][]int, size int, start int, can []int, target int, total int) {
	var (
		res []int
		i   int
	)
	if total == target {
		res = make([]int, len(*route))
		copy(res, *route)
		*results = append(*results, res)
	}
	for i = start; i < size; i++ {
		if i > start && can[i] == can[i-1] {
			continue
		}
		total += can[i]
		if total > target {
			total -= can[i]
			break
		}
		*route = append(*route, can[i])
		backTrack(route, results, size, i+1, can, target, total)
		total -= can[i]
		*route = (*route)[:len(*route)-1]

	}
}

func main() {
	var r []int
	res := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	for _, r = range res {
		fmt.Println(r)
	}
}
