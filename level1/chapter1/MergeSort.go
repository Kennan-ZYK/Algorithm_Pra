package main

import "fmt"

// func mergeSort(arr []int) {
// 	if arr == nil || len(arr) < 2 {
// 		return
// 	}
// 	process(arr, 0, len(arr)-1)
// }

func process(arr []int, left int, right int) {
	if left == right {
		return
	}
	mid := left + ((right - left) >> 1)
	process(arr, left, mid)      //左分治
	process(arr, mid+1, right)   //右分治
	merge(arr, left, mid, right) //分治合并
}

func merge(arr []int, left int, mid int, right int) {
	tmp := make([]int, right-left+1)
	p1, p2 := left, mid+1
	i := 0

	for p1 <= mid && p2 <= right {
		if arr[p1] <= arr[p2] {
			tmp[i] = arr[p1]
			p1++
		} else {
			tmp[i] = arr[p2]
			p2++
		}
		i++
	}
	//右边数组分治结束，左边还未
	for p1 <= mid {
		tmp[i] = arr[p1]
		p1++
		i++
	}
	//左边数组分治结束，右边还未
	for p2 <= right {
		tmp[i] = arr[p2]
		p2++
		i++
	}
	//分治结束，替换原值
	for i := 0; i < len(tmp); i++ {
		arr[left+i] = tmp[i]
	}
}

func main() {
	// var data = []int{3, 5, 75, 34, 8, 31, 44, 76, 87, 25, 73, 11}
	var data = []int{1, 3, 2, 5, 4}
	process(data, 0, len(data)-1)
	fmt.Println(data) //[3 5 8 11 25 31 34 44 73 75 76 87]
}
