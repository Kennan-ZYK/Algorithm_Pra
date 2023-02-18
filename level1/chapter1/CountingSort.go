package main

// 计数排序
func CountingSort(arr []int) {
	ans := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ans[arr[i]]++
	}
}
