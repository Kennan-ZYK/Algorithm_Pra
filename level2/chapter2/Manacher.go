package chapter2

import "math"

// 解决字符串最长回文长度

/*
#a#1#2#1#b#
回文直径 #1#2#1# 7
回文半径 2#1# 4
回文中心 2
[]pArr记录每个位置的回文半径
*/

func manacheString(str string) []byte {

	ch := []byte(str)
	n := len(str)
	ans := make([]byte, 2*n+1)
	j := 0
	for i := 0; i < 2*n+1; i++ {
		if i%2 == 0 {
			ans[i] = '#'
		} else {
			ans[i] = ch[j]
			j++
		}
	}
	return ans
}

func MaxLcpsLenght(s string) int {
	if len(s) == 0 {
		return -1
	}
	// 加入特殊字符
	str := manacheString(s)
	// 记录每个位置的回文半径数组
	pArr := make([]int, len(str))
	// 回文中心，回文右边界
	C, R := -1, -1
	max := math.MinInt64
	for i := 0; i < len(str); i++ {
		// pArr 至少的回文区间
		if R > i {
			// 2*C-i 是 i'位置，R-i是 [i,R]
			//最短的回文半径可能是i对应的i'的回文半径或者i到R的距离
			pArr[i] = min(pArr[2*C-i], R-i)
		} else {
			// 情况1 长度至少1
			pArr[i] = 1
		}

		// 往左右两边扩是否越界，不越界就进行判断
		// 这么写是为了缩短代码
		for i+pArr[i] < len(str) && i-pArr[i] > -1 {
			// 情况1 和 情况2 的 3)
			if str[i+pArr[i]] == str[i-pArr[i]] {
				pArr[i]++
			} else {
				// 情况2 的 1) 2)
				break
			}
		}
		// 查看是否有更往右
		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}
		max = Max(max, pArr[i])
	}

	return max - 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
