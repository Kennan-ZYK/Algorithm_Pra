package main

import (
	"fmt"
)

func main() {
	arr := []int{999, 432, 5, 4326, 2, 27, 7, 72, 4, 2564, 1}
	high := len(arr) - 1
	quicksort(arr, 0, high)
	fmt.Println(arr)
}

func quicksort(arr []int, left, right int) {
	if left >= right {
		return
	}
	//fmt.Println(arr)
	first := arr[left] //第一个元素
	start := left      //记住开始位置
	stop := right      //记住结束位置
	for left != right {
		//找到大于first的值,则一直左移,直到找到小于first的值跳出
		for right > left && arr[right] >= first {
			right-- //指针左移
		}
		//1. 从指针左往右移,比较start 和 arr[0] 大小
		//2. 找到小于first的则start继续右移,
		//3. 直到找到大于first的跳出,
		//4. 该跳出位置的start等待与 从右往左移动的找到的小于first的交换
		for left < right && arr[left] <= first {
			left++ //继续右移
		}
		if right > left { //right >left 表示,左右指针还未重合,数据未交换完毕,
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	//当左右指针相等时,代表查找到中间位置了，
	//该中间位置左侧所有数据均小于 该位置右侧数据
	//把第一个first与该位置值交换,则 该位置左侧所有值,都小于该中间位置值,右侧所有值都大于该位置的值
	arr[start], arr[right] = arr[right], first
	quicksort(arr, start, left-1)
	quicksort(arr, right+1, stop)
}
