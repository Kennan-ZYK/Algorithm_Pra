// package chapter1
package main

import "fmt"

/*题目：
给定一个有序数组arr，代表数轴上从左到右有n个点arr[0]、arr[1]...arr[n-1]。
给定一个正数L，代表一根长度为L的绳子，求绳子最多能覆盖其中的几个点。
*/

/*思路:
用滑动窗口的思路，从第一个元素开始，窗口长度最大为5
绳子左侧压中 数组arr的点
*/
// 时间复杂度：O(N)

func MaxPoint(arr []int, L int) int {
	res := 1
	for i := 0; i < len(arr); i++ {
		tmp := Window(arr, i, L)
		res = Max(res, tmp)
	}
	return res
}

func Window(arr []int, start int, L int) int {
	res := 1
	left, right := start, start+1
	for right <= L && L+arr[left] >= arr[right] {
		res++
		right++
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{0, 13, 24, 35, 45, 67, 72, 87}
	L := 6
	fmt.Println(MaxPoint(arr, L))
}
