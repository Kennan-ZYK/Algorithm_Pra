// package chapter1
package main

import "fmt"

/*题目：
给定一个有序数组arr，代表数轴上从左到右有n个点arr[0]、arr[1]...arr[n-1]。
给定一个正数L，代表一根长度为L的绳子，求绳子最多能覆盖其中的几个点。
*/

/*思路:
用绳子的右端点去压 数组arr的点，然后用二分查找 大于等于L-arr[i]最左侧的点。
再用 i-最左侧点的下标，即为覆盖了多少点
*/
// 时间复杂度：N*logn  （N个点，然后二分logn）

func MaxPoint(arr []int, L int) int {
	res := 1
	for i := 0; i < len(arr); i++ {
		value := arr[i] - L
		tmp := leftnear(arr, value, i)
		res = Max(res, i-tmp+1)
	}
	return res
}

func leftnear(arr []int, value int, i int) int {
	index := 0
	left, right := 0, i

	for left <= right {
		mid := left + ((right - left) >> 1)
		if arr[mid] >= value {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
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
