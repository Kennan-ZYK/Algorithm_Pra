package main

import "fmt"

func QuickSort(array []int) {
	SortDetail(array, 0, len(array)-1)
}

//l,r分别为数组的最左、最右下标
func SortDetail(array []int, l, r int) {
	//递归一定要有截止条件
	if l >= r {
		return
	}
	point := Partition(array, l, r)
	SortDetail(array, l, point-1)
	SortDetail(array, point+1, r)
}

//分区函数
func Partition(array []int, l, r int) int {
	//pivot为分区点，这里选择数组的最后一个元素作为分区点
	point := array[r]
	//i元素为数组最左元素的下标
	i := l
	for j := l; j < r; j++ {
		if array[j] < point {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	//上面for循环只是到r-1,这里处理最后一个元素pivot
	array[i], array[r] = array[r], array[i]
	return i
}

func main() {
	array := []int{3, 1, 2, 4, 232, 5, 3, 5, 6, 22, 6, 47, 1, 2, 2, 2, 66, 23123, 123, 4, 23, 66, 1}
	QuickSort(array)
	// for i, v := range array {
	// 	fmt.Printf("下标为%d的值为%d", i, v)
	// 	fmt.Println()
	// }
	fmt.Println(array)
}
