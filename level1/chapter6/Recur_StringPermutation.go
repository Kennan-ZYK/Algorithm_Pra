package main

import (
	"fmt"
)

/*打印一个字符串的全部排列，要求不要出现重复的排列*/

// 思路：让每个字符在每个位置都待过，然后在去重。

func Permutation(str string) *[]string {
	res := []string{}
	if len(str) == 0 {
		return &res
	}
	ch := []byte(str)
	ans := process(ch, 0, &res)
	return ans
}

func process(str []byte, i int, res *[]string) *[]string {
	if i == len(str) {
		(*res) = append((*res), string(str))
		return res
	}
	// 判定去重
	visit := [26]bool{}
	for j := i; j < len(str); j++ {
		if !visit[str[j]-'a'] { // 可优化常数时间
			visit[str[j]-'a'] = true
			swap(str, i, j)
			process(str, i+1, res)
			// 回溯
			swap(str, i, j)
		}
	}
	return res
}

func swap(str []byte, i, j int) {
	tmp := str[i]
	str[i] = str[j]
	str[j] = tmp

}

func main() {
	ans := Permutation("abc")
	fmt.Println(ans)
}
