package main

/**
 * 快速排序1.0
 * 每次可以保证一个数是排完的
 * 以最后一个数num为基准，做到 [<num]  [num]   [>num],
 * 找到>num区间的第一位与最后一位数num交换，即完成一次
 * O(N^2),最坏情况是[1,2,3,4,5] 每次都要遍历完
 */
import "fmt"

func quickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	pos := partition(arr, left, right)
	quickSort(arr, left, pos-1)
	quickSort(arr, pos+1, right)

}

func partition(arr []int, left int, right int) int {
	//分区的最后一个元素作为分区点
	pos := right
	//设置左区间
	less := left - 1

	for left < right {
		if arr[left] < arr[pos] {
			less++
			arr[left], arr[less] = arr[less], arr[left]
			left++
		} else {
			left++
		}
	}
	//交换后，表示一次快排结束
	//这样就可以保证[left,less+1]是小于等于pos的
	arr[pos], arr[less+1] = arr[less+1], arr[pos]

	return less + 1
}

func main() {
	array := []int{3, 2, 26, 3, 66, 22, 1}
	quickSort(array, 0, len(array)-1)
	// for i, v := range array {
	// 	fmt.Printf("下标为%d的值为%d", i, v)
	// 	fmt.Println()
	// }
	fmt.Println(array)
}
