package main

import (
	"fmt"
)

/*打印一个字符串的全部排列，要求不要出现重复的排列*/

// 思路：让每个字符在每个位置都待过，然后在去重。
var res [][]int

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res = [][]int{}
	ans := process(nums, 0)
	return ans
}

func process(nums []int, i int) [][]int {
	if i == len(nums) {
		// 这里不能 res = append(res,nums),因为用nums每次都是同一块地址，在swap会导致res内的nums也改变
		// 所以用copy复制再放入，这样地址就改变了不会影响到了,不能用变量结果一样，只能用copy
		p := make([]int, len(nums))
		copy(p, nums)
		res = append(res, p)
	}
	for j := i; j < len(nums); j++ {
		swap(nums, i, j)

		process(nums, i+1)
		swap(nums, i, j)
	}
	return res
}

// 打印不重复
func process2(nums []int, i int) [][]int {
	if len(nums) == i {
		p := make([]int, len(nums))
		copy(p, nums)
		res = append(res, p)
		return res
	}
	// -1 - 9 21个数
	visit := [21]bool{}
	for j := i; j < len(nums); j++ {
		if !visit[nums[j]+10] {
			visit[nums[j]+10] = true
			swap(nums, i, j)
			process(nums, i+1)
			// 回溯
			swap(nums, i, j)
		}
	}
	return res
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
func main() {
	ans := permute([]int{1, 3, 4})
	fmt.Println(ans)
}
