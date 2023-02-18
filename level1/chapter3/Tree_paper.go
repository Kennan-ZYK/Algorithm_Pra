package main

import "fmt"

// 折纸问题：每次向上对折，打印出折纸后的情况
/*
第一次折：1凹
第二次折：2凹 1凹 2凸
第三次折：3凹 2凹 3凸 1凹 3凹 2凸 3凸
...
以此类推
得出规律：第一次之后，每对折一次，就会在原有的凹上下各加凹凸
可以看作二叉树，左子树为凹，右子树为凸
*/
// 	 	凹
//   凹    凸
// 凹  凸 凹 凸

// 参数：
// 第i层 开始
// n 代表对折次数
// down = 凹 (true)  凸(false)
func paper(n int) {
	process(1, n, true)
}
func process(i int, n int, down bool) {
	if i > n {
		return
	}
	process(i+1, n, true)
	if down == true {
		fmt.Println("凹")
	} else {
		fmt.Println("凸")
	}
	process(i+1, n, false)
}

func main() {
	paper(2)
}
