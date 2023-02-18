package chapter1

/*
str	 a  b  b s t a b b z
next[0,-1, 0,0,0,0,1,2,3]

*/

// 判断字符串1是否包含字符串2 ，返回开始位置
// next数组使保存前缀最长字符串个数，如 next[i]保存 0~i-1中有相等字符串的最大长度
// 时间复杂度O(M+N) -> O(N) 因为M的长度不可能大于N
func getIndexOf(str, model string) int {
	if len(model) > len(str) || len(model) == 0 {
		return -1
	}
	str1 := []byte(str)
	str2 := []byte(model)

	next := getNext(str2) //O(M)
	i1, i2 := 0, 0
	// O(N)
	for i1 < len(str1) && i2 < len(str2) {
		// 相等同时++
		if str1[i1] == str2[i2] {
			i1++
			i2++
		} else if next[i2] == -1 { //也可写(i2 ==0). 跳转后全部可能性还不匹配，只能i1下一个字符了
			i1++
		} else { // 跳转到上一个最长字符串的后缀点
			i2 = next[i2]
		}
	}
	// 全匹配返回起始下标
	if i2 == len(str2) {
		return i1 - i2
	} else {
		// i1越界
		return -1
	}
}

func getNext(ch []byte) []int {
	if len(ch) == 0 {
		return []int{-1}
	}
	next := make([]int, len(ch))
	// 人为调值 0 和 1 位置
	next[0] = -1
	next[1] = 0
	i := 2
	cn := 0
	for i < len(ch) {
		if ch[i-1] == ch[cn] {
			next[i] = cn + 1
			i++
		} else if cn > 0 { //和i-1位置字符不匹配，跳到cn位置
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}
