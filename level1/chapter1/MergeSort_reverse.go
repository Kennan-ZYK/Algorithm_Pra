package main

import "fmt"

/*
题目:逆序对问题在一个数组中，左边的数如果比右边的数大，则折两个数构成一个逆序对，请打印所有逆序对。
*/

func process(arr []int, left int, right int) int {
	if left == right {
		return 0
	}
	mid := left + ((right - left) >> 1)
	return process(arr, left, mid) + process(arr, mid+1, right) + merge(arr, left, mid, right)

}

func merge(arr []int, left int, mid int, right int) int {
	tmp := make([]int, right-left+1)
	p1, p2 := mid, right
	i := len(tmp) - 1
	res := 0
	for p1 >= left && p2 >= mid+1 {
		if arr[p1] > arr[p2] {
			// 找到右组第一个比左组小的数，则当前满足要求的逆序对个数为 (p2 - (mid + 1) + 1) 即是 (pR - mid)
			res += p2 - mid
			tmp[i] = arr[p1]
			p1--
		} else {
			res += 0
			tmp[i] = arr[p2]
			p2--
		}
		i--
	}

	for p1 >= left {
		tmp[i] = arr[p1]
		p1--
		i--

	}
	//左边数组分治结束，右边还未
	for p2 >= mid+1 {
		tmp[i] = arr[p2]
		p2--
		i--
	}
	for i := 0; i < len(tmp); i++ {
		arr[left+i] = tmp[i]
	}
	return res
}

func main() {
	// var data = []int{3, 5, 75, 34, 8, 31, 44, 76, 87, 25, 73, 11}
	var data = []int{3, 5, 2, 1, 0, 4, 9}
	a := process(data, 0, len(data)-1)
	fmt.Println(data, a) //[3 5 8 11 25 31 34 44 73 75 76 87]
}
