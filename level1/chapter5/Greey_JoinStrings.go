package main

import (
	"fmt"
	"sort"
)

/*
题目:拼接字符串
给你一组字符串，要求拼接后的字符串最大。

思路：按字符串之间相加来比较大小进行排序
*/

type StringSlice []string

func (x StringSlice) Len() int           { return len(x) }
func (x StringSlice) Less(i, j int) bool { return x[i]+x[j] > x[j]+x[i] }
func (x StringSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func joinStrings(strings []string) string {
	if len(strings) == 0 {
		return ""
	}
	sort.Sort(StringSlice(strings))
	res := ""
	for i := 0; i < len(strings); i++ {
		res += strings[i]
	}
	return res
}

func main() {
	a := []string{"qew", "zxc", "ggd", "zz"}
	res := joinStrings(a)
	fmt.Println(res)
}
