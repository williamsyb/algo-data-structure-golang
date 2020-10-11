package main

import "fmt"

/*
438. 找到字符串中所有字母异位词
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：
字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。
示例 1:
输入:
s: "cbaebabacd" p: "abc"
输出:
[0, 6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 示例 2:

输入:
s: "abab" p: "ab"
输出:
[0, 1, 2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
*/

func findAnagrams(s string, p string) []int {
	var (
		need   map[byte]int
		window map[byte]int
		//size int
		res   []int
		char  byte
		left  int
		right int
		i     int
		ok    bool
		valid int
	)
	res = []int{}
	need = make(map[byte]int)
	window = make(map[byte]int)
	for i = 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right = 0, 0
	for right < len(s) {
		char = s[right]
		if _, ok = need[char]; ok {
			window[char]++
			if window[char] <= need[char] {
				valid++
			}
		}
		right++
		for valid == len(p) {
			if right-left == len(p) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok = need[d]; ok {

				if window[d] <= need[d] {
					valid--
				}
				window[d]--
			}

		}

	}
	return res
}

func findAnagramsV2(s string, p string) []int {
	//不使用模板的方法
	/*
		执行用时：4 ms, 在所有 Go 提交中击败了97.83%的用户
		内存消耗：5 MB, 在所有 Go 提交中击败了99.62%的用户
	*/
	if len(p) > len(s) {
		return nil
	}
	var res []int
	mat01 := [26]int{}
	mat02 := [26]int{}
	for index := range p {
		mat01[s[index]-'a']++
		mat02[p[index]-'a']++
	}
	for i := 0; i <= len(s)-len(p); i++ {
		if mat01 == mat02 {
			res = append(res, i)

		}
		mat01[s[i]-'a']--
		if i+len(p) < len(s) {
			mat01[s[i+len(p)]-'a']++
		}

	}
	return res
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagramsV2("abab", "ab"))
}
