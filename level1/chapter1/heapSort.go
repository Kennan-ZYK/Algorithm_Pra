package main

import "fmt"

// 堆排序

func heapSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr); i++ { //O(N)
		heapInsert_Big(&arr, i) //O(logn)
	}
	//去除头元素，换头
	heapSize := len(arr)
	heapSize--
	swap(arr, 0, heapSize)
	for heapSize > 0 { //O(N)
		heapify(arr, 0, heapSize) //O(logn)
		heapSize--
		swap(arr, 0, heapSize) //O(1)
	}
}

func heapInsert_Big(arr *[]int, index int) {
	//处于index位置，与其父节点比较，是否需要交换
	for (*arr)[index] > (*arr)[(index-1)/2] {
		swap((*arr), index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func heapInsert_Small(arr []int, index int) {
	for arr[index] < arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func heapify(arr []int, index int, heapSize int) {
	//left为index的左子节点
	left := index*2 + 1
	max := 0
	for left < heapSize {
		// 比较两个子结点大小，若没有右结点，max=left
		if arr[left+1] < heapSize && arr[left+1] > arr[left] {
			max = left + 1
		} else {
			max = left
		}
		// max孩子与父节点比较，若儿子小于父亲小于，则调整到位
		if arr[max] < arr[index] {
			max = index
		}
		// 调整到位了，直接break
		if max == index {
			break
		}
		//否则交换
		swap(arr, max, index)
		index = max
		left = index*2 + 1

	}
}
func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	var index []int = []int{4, 2, 41, 2135, 7, 65, 8}
	heapSort(index)
	fmt.Println(index)
}
