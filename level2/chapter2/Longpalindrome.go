package main

import (
	"fmt"
	"math"
)

// 最长回文字符串

/*
暴力思路：在每个字符前后都加入一个字符(可以避免奇偶数情况)，然后从左往右开始遍历字符串，
		 每个字符以自己为中心，向两边扩，记录最长的相同字符串长度,最后除于2就是答案
举例:
# a # b # b # a # g # c # b # b # a #
1 3 1 3 9 3 1 3 1 3 1 3 1 3 5 3 1 3 1

9 / 2 = 4 (abba)
*/
// O(N)
func longpalindrome(str string) []int {

	n := len(str)
	tmp := make([]byte, n+n+1)
	// 插入字符
	for i := 0; i < n+n+1; i = i + 2 {
		tmp[i] = '#'
	}
	for i, j := 1, 0; i < n+n+1; i, j = i+2, j+1 {
		tmp[i] = str[j]
	}
	// fmt.Println(string(tmp))
	p, q := 0, 0

	ans := make([]int, n+n+1)
	for i := 0; i < n+n+1; i++ {
		t := i
		p = t - 1
		q = t + 1
		cnt := 1
		// 边界情况
		if p < 0 || q > n+n {
			ans[i] = cnt
			continue
		}
		// fmt.Println("ASdsad")
		for p >= 0 && q <= n+n {
			if tmp[p] != tmp[q] {
				ans[i] = cnt
				break
			} else {
				cnt = cnt + 2
				ans[i] = cnt
			}
			p--
			q++
		}
		// 走else，然后p--,q++越界，cnt要重新初始
	}
	// fmt.Println(ans)
	return ans
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	st := "abaaxjsowlaabaa"
	ans := longpalindrome(st)
	max := math.MinInt64
	for i := 0; i < len(ans); i++ {
		max = Max(max, ans[i])
	}
	fmt.Println(max / 2)
}
