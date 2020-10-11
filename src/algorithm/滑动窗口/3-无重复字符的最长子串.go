package main

import "fmt"

/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func lengthOfLongestSubstring(s string) int {
	var (
		left   int
		right  int
		res    int
		char   byte
		window map[byte]int
		ok     bool
		maxLen int
	)
	window = make(map[byte]int)
	res = 0
	maxLen = 0
	left, right = 0, 0
	for right < len(s) {
		char = s[right]

		if _, ok = window[char]; !ok {
			window[char]++
			res++
			if res > maxLen {
				maxLen = res
			}
			right++
		} else {
			for left < right {
				d := s[left]
				if char != d {
					left++
					delete(window, d)
					res--
				} else {
					res--
					left++
					delete(window, d)
					break
				}
			}
		}
	}
	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	//maxLen := 0

	for right < len(s) {
		char := s[right]
		window[char]++
		right++
		for window[char] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring2("tmmzuxt"))
}
