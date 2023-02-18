package main

import (
	"fmt"
	"math"
)

/*
基数排序
*/

func RadixSort(arr []int) {
	if arr == nil && len(arr) < 2 {
		return
	}
	radixSort(arr, 0, len(arr)-1, maxbits(arr))
}

//求最高的位数
func maxbits(arr []int) int {
	// max := ^int(^uint(0) >> 1)
	// 系统最小值
	max := math.MinInt64
	//得出最大值
	for i := 0; i < len(arr); i++ {
		max = Max(max, arr[i])
	}
	// 求最大有几位数
	res := 0
	for max != 0 {
		res++
		max /= 10
	}
	return res

}

//dight 十进制数上最大位数 (131 dight=3)
func radixSort(arr []int, L int, R int, dight int) {
	radix := 10
	bucket := make([]int, R-L+1)
	for d := 1; d <= dight; d++ { //有几位就进出几次
		//10个空间
		// count[0] 当前位(d位)是0的数字有几个
		// count[1] 当前位(d位)是(0，1)的数字有几个
		// count[2] 当前位(d位)是(0,1,2)的数字有几个
		// count[i] 当前位(d位)是(i)的数字有几个
		count := make([]int, radix)

		//判断各个位上的数分别有多少（词频）
		for i := L; i <= R; i++ {
			j := getDigit(arr[i], d)
			count[j]++
		}
		//前缀和，作用是可以快速看出有多少个 小于等于i的数
		for i := 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}
		//从右往左开始，因为右边是最后放入，就定好一个进制数的边界
		// 不需要关注其他位置上的数，因为彼此之间互不影响，只是说前面有几个数小于i的数
		for i := R; i >= L; i-- {
			j := getDigit(arr[i], d)
			bucket[count[j]-1] = arr[i]
			count[j]--
		}
		// 出桶
		for i, j := L, 0; i <= R; i, j = i+1, j+1 {
			arr[i] = bucket[j]
		}
	}
}

//获取某位上的数
func getDigit(x int, d int) int {
	return (x / int(math.Pow(10, float64(d-1)))) % 10

}

func Max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	a := []int{13, 21, 11, 52, 62}
	RadixSort(a)
	fmt.Println(a)
}
