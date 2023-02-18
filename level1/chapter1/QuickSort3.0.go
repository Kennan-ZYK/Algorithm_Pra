/**
 * 快速排序3.0
 * 1.0和2.0版本的时间复杂度都是O(N^2)
 * 3.0版本时间复杂度是O(NlogN)
 * 随机设置point，好情况差情况都是等概率事件，T(N) = 2T(N/2) + O(N),T(N) = T(N/3) + T(2N/3)+ O(N)...等情况
 * 每种情况发生概率都为1/N，根据数学期望最后算出为O(NlogN)
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//快排
func QuickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	//随机选择一位作为point，
	rand.Seed(time.Now().UnixNano())
	swap(arr, left+int(rand.Float64()*float64(right-left+1)), right)
	pos := partition(arr, left, right)
	QuickSort(arr, left, pos[0])
	QuickSort(arr, pos[1], right)
}

//分化,告知左右边界
func partition(arr []int, left int, right int) []int {
	//定义大于小于区间
	less, more := left-1, right+1
	pos := arr[right]
	for left < more {
		if arr[left] < pos {
			//扩充小于区间，再交换
			less++
			swap(arr, left, less)
			left++
		} else if arr[left] > pos {
			//扩充大于区间
			more--
			swap(arr, left, more)
			//这里为什么不left++呢？
			//因为我们是把more区间前一个元素(more--)进行交换，还需在判断
		} else {
			left++
		}
	}
	// swap(arr, more, right) 不需要，因为右边界是right+1
	return []int{less, more}
}

func swap(arr []int, num1 int, num2 int) {
	if num1 == num2 {
		return
	}
	arr[num1] = arr[num1] ^ arr[num2]
	arr[num2] = arr[num1] ^ arr[num2]
	arr[num1] = arr[num1] ^ arr[num2]

}

func main() {
	arr := []int{2, 1999, 432, 5, 43, 4, 26, 2, 5, 7, 7, 27, 7, 72, 4, 2564, 1}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
