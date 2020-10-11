package main

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

func longestPalindrome(s string) string {
	/*
			时间复杂度 O(N^2)，空间复杂度 O(1)。

			值得一提的是，这个问题可以用动态规划方法解决，时间复杂度一样，
			但是空间复杂度至少要 O(N^2) 来存储 DP table。这道题是少有的动态规划非最优解法的问题。
			用双指针。

			寻找回文串的问题核心思想是：从中间开始向两边扩散来判断回文串。对于最长回文子串，就是这个意思：

			for 0 <= i < len(s):
			    找到以 s[i] 为中心的回文串
			    更新答案
			但是呢，我们刚才也说了，回文串的长度可能是奇数也可能是偶数，如果是abba这种情况，没有一个中心字符，上面的算法就没辙了。
		    所以我们可以修改一下：

			for 0 <= i < len(s):
			    找到以 s[i] 为中心的回文串
			    找到以 s[i] 和 s[i+1] 为中心的回文串
			    更新答案

	*/
	var (
		sByte []byte
		s1    []byte
		s2    []byte
		i     int
		res   []byte
	)
	sByte = []byte(s)
	s1 = []byte{}
	res = []byte{}
	s2 = []byte{}
	for i = 0; i < len(sByte); i++ {
		s1 = palindrome(sByte, i, i)
		s2 = palindrome(sByte, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return string(res)
}

func palindrome(s []byte, left int, right int) []byte {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
