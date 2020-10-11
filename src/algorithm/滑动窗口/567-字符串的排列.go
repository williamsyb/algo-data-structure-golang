package main

import "fmt"

/*
567. 字符串的排列
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").


示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False
*/

func checkInclusion(s1 string, s2 string) bool {
	// 套用labuladong模板方法，模板的方法最通用，但是效率不总是最高的，模板里的窗口是动态变化的，需要借助其他信息来判断何时变化，
	// 但是本题很明显可以看出窗口大小是不变的，使用模板将导致效率低下。
	// 因此，需要在掌握模板的基础上灵活的去使用滑动窗口
	var (
		left   int
		right  int
		need   map[byte]int
		window map[byte]int
		size   int
		s1Byte []byte
		s2Byte []byte
		i      int
		ok     bool
		char   byte
		valid  int
	)
	need = make(map[byte]int)
	window = make(map[byte]int)
	left = 0
	right = 0
	size = 0
	valid = 0
	s1Byte = []byte(s1)
	s2Byte = []byte(s2)
	for i = 0; i < len(s1Byte); i++ {
		need[s1Byte[i]]++
	}
	size = len(s1Byte)
	for right < len(s2Byte) {
		char = s2Byte[right]
		if _, ok = need[char]; ok {
			window[char]++
			if window[char] <= need[char] {
				valid++
			}
		}
		right++
		for valid == len(s1Byte) {
			if right-left == size {
				return true
			}
			d := s2Byte[left]
			left++
			if _, ok = need[d]; ok {

				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

func checkInclusionV2(s1, s2 string) bool {
	// leetcode 上GO 100%击败的代码，维护一个长度固定的窗口就ok
	/*
		用数组代替map可以比较
		比较相同长度的s1和s2的子串对应的数组是否相等来判断是否为子串
		s2的子串通过滑动窗口来移动
	*/
	if len(s1) > len(s2) {
		return false
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if cnt1 == cnt2 {
			return true
		}
		cnt2[s2[i]-'a']--
		cnt2[s2[i+len(s1)]-'a']++
	}
	return cnt1 == cnt2
}

func main() {
	fmt.Println(checkInclusion("abcdxabcde", "abcdeabcdx"))
	fmt.Println(checkInclusionV2("ab", "eidboaoo"))
}
