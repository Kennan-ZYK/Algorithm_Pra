package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
arr里都是正数，没有重复值，每一个值代表一种货币，每一种都可以用无限张
最终要找零钱数是aim，
找零方法数返回

暴力递归-表结构-斜率优化
斜率优化：填表时，有枚举过程，观察能不能通过临近的位置，代替枚举过程
*/

// 暴力递归
func way1(arr []int, aim int) int {
	if len(arr) == 0 {
		return -1
	}
	return process(arr, 0, aim)
}

func process(arr []int, index, aim int) int {
	if index == len(arr) {
		if aim == 0 {
			return 1
		} else {
			return 0
		}
	}
	way := 0
	for zhang := 0; zhang*arr[index] <= aim; zhang++ {
		way += process(arr, index+1, aim-arr[index]*zhang)
	}

	return way
}

// 表结构
func way2(arr []int, aim int) int {
	if len(arr) == 0 {
		return 0
	}
	// 0~len(arr)
	dp := make([][]int, len(arr)+1)
	for i := 0; i < len(arr)+1; i++ {
		dp[i] = make([]int, aim+1)
	}
	dp[len(arr)][0] = 1

	// CoinWay图，index 等于倒数第二行开始
	for index := len(arr) - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			way := 0
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				way += dp[index+1][rest-arr[index]*zhang]
			}
			dp[index][rest] = way
		}
	}
	// fmt.Println(dp)
	return dp[0][aim]
}

func way3(arr []int, aim int) int {
	if len(arr) == 0 {
		return 0
	}

	dp := make([][]int, len(arr)+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[len(arr)][0] = 1

	for index := len(arr) - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				dp[index][rest] += dp[index+1][rest-arr[index]*zhang]
				if rest-arr[index] >= 0 {
					// 找最近的，所以只需要一张rest-arr[index]*1
					return dp[index][rest-arr[index]]
				}
			}
		}
	}
	return dp[0][aim]

}

// func main() {
// 	arr := []int{3, 5, 1, 2}
// 	aim := 5
// 	fmt.Printf("way2(arr, aim): %v\n", way2(arr, aim))
// 	fmt.Printf("way1(arr, aim): %v\n", way1(arr, aim))
// }

// 对数器
func generateRandomArry(lenght int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, int(rand.Float32()*float32(lenght))+1)
	for i := 0; i < len(arr); i++ {
		// +1 是因为rand.Float32[0,x),得加1才能得到x
		arr[i] = int(rand.Float32()*float32(max)) + 1
	}
	return arr
}

func main() {
	lenght, max := 10, 10
	testime := 100
	for i := 0; i < testime; i++ {
		arr := generateRandomArry(lenght, max)
		aim := int(rand.Float32()*float32(max)*3) + max
		// fmt.Println(way1(arr, aim))
		if way1(arr, aim) != way2(arr, aim) && way2(arr, aim) != way3(arr, aim) {
			// fmt.Println(CoinMin2(arr, aim), CoinMin3(arr, aim))
			fmt.Println("xxx")
			break
		}
	}
	fmt.Println("right")
}
