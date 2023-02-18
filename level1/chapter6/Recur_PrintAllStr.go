package main

import (
	"fmt"
	"strings"
)

/*打印一个字符串的全部子序列，包含空字符串*/

// 也等于求一个字符串的子集

// 思路 类似折纸问题
func PrintStr(str string) {
	ch := strings.Split(str, "")
	process(ch, 0)
}

func process(ch []string, i int) {
	if len(ch) == i {
		print(ch)
		return
	}
	//
	process(ch, i+1)
	tmp := ch[i]
	ch[i] = ""
	process(ch, i+1)
	// 回滚
	ch[i] = tmp
}

func print(str []string) {
	ans := ""
	for _, v := range str {
		ans += v
	}
	fmt.Printf("%s\n", ans)

}

func main() {
	PrintStr("abv")
}
