package main

import "fmt"

// 在一个数组中， 对于每个数num，求有多少个后面的数 * 2 依然<num，求总个数
func biggerTwice(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	return process(arr, 0, len(arr)-1)
}
func process(arr []int, left int, right int) int {
	if left == right {
		return 0
	}

	mid := left + ((right - left) >> 1)
	return process(arr, left, mid) + process(arr, mid+1, right) + merge(arr, left, mid, right)
}

func merge(arr []int, left int, mid int, right int) int {
	//先计算有几个
	ans := 0
	windowR := mid + 1
	for j := left; j <= mid; j++ {
		for windowR <= right && (arr[j] > arr[windowR]*2) {
			fmt.Printf("%d,%d\n", arr[j], arr[windowR])
			windowR++
		}
		ans += windowR - mid - 1
	}

	tmp := make([]int, right-left+1)
	p1, p2 := left, mid+1
	i := 0
	for p1 <= mid && p2 <= right {
		if arr[p1] < arr[p2] {
			tmp[i] = arr[p1]
			p1++
		} else {
			tmp[i] = arr[p2]
			p2++
		}
		i++
	}
	for p1 <= mid {
		tmp[i] = arr[p1]
		p1++
		i++
	}

	for p2 <= right {
		tmp[i] = arr[p2]
		p2++
		i++
	}
	// fmt.Println(len(tmp))
	for i := 0; i < len(tmp); i++ {
		arr[left+i] = tmp[i]
	}
	return ans
}
func main() {
	var data = []int{2, 3, 5, 1}
	// var data = []int{3, 5, 2, 1, 0, 4, 9}
	a := biggerTwice(data)
	fmt.Println(data, a) //[3 5 8 11 25 31 34 44 73 75 76 87]
}
