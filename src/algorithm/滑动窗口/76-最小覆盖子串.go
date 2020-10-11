package main

import (
	"fmt"
)

/*
76. 最小覆盖子串
困难

给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。

示例：
输入：S = "ADOBECODEBANC", T = "ABC"
输出："BANC"


提示：
如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。

*/

const MAX_INT = int(^uint(0) >> 1)

//解析：https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485141&idx=1&sn=0e4583ad935e76e9a3f6793792e60734&chksm=9bd7f8ddaca071cbb7570b2433290e5e2628d20473022a5517271de6d6e50783961bebc3dd3b&scene=21#wechat_redirect

func minWindow(s string, t string) string {
	var (
		need   map[byte]int
		window map[byte]int
		left   int
		right  int
		valid  int
		start  int
		size   int
		char   byte
		sByte  []byte
		tByte  []byte
		ok     bool
	)
	left = 0
	right = 0
	valid = 0
	// 记录最小覆盖子串的起始索引及长度
	start = 0
	size = MAX_INT
	sByte = []byte(s)
	need = make(map[byte]int)
	window = make(map[byte]int)
	tByte = []byte(t)

	//初始化目标need哈希表
	for i := 0; i < len(tByte); i++ {
		need[tByte[i]]++
	}

	for right < len(sByte) {
		//char是将移入窗口的字符
		char = sByte[right]
		// 右移窗口
		right++
		if _, ok = need[char]; ok {
			// 进行窗口内数据的一系列更新
			window[char]++
			if window[char] == need[char] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < size {
				start = left
				size = right - left
			}
			// d 是将移出窗口的字符
			d := sByte[left]
			// 左移窗口
			left++
			if _, ok = need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if size == MAX_INT {
		return ""
	} else {
		return string(sByte[start : start+size])
	}
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
