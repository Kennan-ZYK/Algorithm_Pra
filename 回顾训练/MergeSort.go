package main

import "fmt"

func process(nums []int, left int, right int) {
	if len(nums) == 0 {
		return
	}
	if left == right {
		return
	}
	mid := left + (right-left)>>1
	// 左分治
	process(nums, left, mid)
	// 右分治
	process(nums, mid+1, right)

	MergeSort(nums, left, mid, right)
}

func MergeSort(nums []int, left int, mid int, right int) []int {
	tmp := make([]int, right-left+1)
	i := 0
	p1, p2 := left, mid+1

	for p1 <= mid && p2 <= right {
		if nums[p1] >= nums[p2] {
			tmp[i] = nums[p2]
			p2++
			i++
		} else {
			tmp[i] = nums[p1]
			p1++
			i++
		}
	}
	for p1 <= mid {
		tmp[i] = nums[p1]
		p1++
		i++
	}
	for p2 <= right {
		tmp[i] = nums[p2]
		p2++
		i++
	}

	for i := 0; i < len(tmp); i++ {
		nums[left+i] = tmp[i]
	}
	return nums
}

func main() {
	// var data = []int{3, 5, 75, 34, 8, 31, 44, 76, 87, 25, 73, 11}
	var data = []int{1, 3, 2, 5, 4}
	process(data, 0, len(data)-1)
	fmt.Println(data) //[3 5 8 11 25 31 34 44 73 75 76 87]
}
