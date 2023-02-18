package main

import "fmt"

func MaxGlobal(nums []int) int {
	return process(nums, 0, len(nums)-1)
}

func process(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := left + (right-left)>>1
	leftMax := process(nums, left, mid)
	rightMax := process(nums, mid+1, right)
	return max(leftMax, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 3, 4, 5, 2}
	ans := MaxGlobal(nums)
	fmt.Println(ans)
}
