/**
 * 快速排序2.0
 * 在1.0的基础上将最右侧数的一组放在中间，左侧都比其小，右侧都比其大
 * 比1.0快一点，因为一次搞定一批数
 * O(N^2),因为1 2 3 4 5,每次都是都划分一个值，n+n-1+n-2+...+2+1, 等差数列
 */

package main

import "fmt"

//快排
func QuickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	//随机将数组中一个数放在数组的最右侧

	pos := partition(arr, left, right)
	QuickSort(arr, left, pos[0])
	QuickSort(arr, pos[1], right)
}

//分化
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
	arr := []int{1, 2, 3, 4, 6, 7, 6, 5, 6}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
