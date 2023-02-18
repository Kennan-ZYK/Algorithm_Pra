package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 题目: 给一个数组arr ，每个arr[i]代表一枚硬币， 给定一个aim，组成aim最少硬币数目（有限硬币）

func CoinMin1(arr []int, aim int) int {
	return process1(arr, 0, aim)
}

// arr 钱币数
// index 起始选金币位置
// aim 目标
// 暴力
func process1(arr []int, index, rest int) int {
	// 要的钱超过aim目标，导致rest<0 或者 已经没有硬币了，
	if rest < 0 || index == len(arr) {
		return -1
	}

	if rest == 0 {
		return 0
	}
	// 不要当前硬币
	p1 := process1(arr, index+1, rest)
	// 要当前硬币 1+p2
	p2 := 1 + process1(arr, index, rest-arr[index])

	// 没有方法可以组成aim
	if p1 == -1 && p2 == -1 {
		return -1
	} else {
		// p1 == -1 ,p2 != -1
		if p1 == -1 {
			return p2
		}
		// p2==-1, p1 != -1
		if p2 == -1 {
			return p1
		}
		return Min(p1, p2)
	}

}

// 定义dp [index+1][aim+1], 0 ,aim 都为终止
// index 行， aim 为列
func CoinMin2(arr []int, aim int) int {
	var dp [][]int
	dp = make([][]int, len(arr)+1)

	// 初始化
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	for i := 0; i <= len(arr); i++ {
		for j := 0; j <= aim; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	return process2(arr, 0, aim, dp)
}

// 记忆搜索
func process2(arr []int, index int, rest int, dp [][]int) int {
	if rest < 0 {
		return -1
	}
	// 证明记录过
	if dp[index][rest] != math.MaxInt64 {
		return dp[index][rest]
	}
	if rest == 0 {
		dp[index][rest] = 0
		return 0
	} else if index == len(arr) {
		dp[index][rest] = -1
		return -1
	}
	// 不要当前硬币
	p1 := process2(arr, index+1, rest, dp)
	// 要当前硬币
	p2 := 1 + process2(arr, index, rest-arr[index], dp)
	if p1 == -1 && p2 == -1 {
		dp[index][rest] = -1
	} else {
		if p1 == -1 {
			dp[index][rest] = p2
		} else if p2 == -1 {
			dp[index][rest] = p1
		} else {
			dp[index][rest] = Min(p1, p2)
		}
	}
	fmt.Println(dp)

	return dp[index][rest]

}

func CoinMin3(arr []int, aim int) int {
	var dp [][]int
	dp = make([][]int, len(arr)+1)
	// 初始化
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}

	for i := 0; i < len(arr)+1; i++ {
		dp[i][0] = 0
	}

	for j := 1; j < aim+1; j++ {
		dp[len(arr)][j] = -1
	}
	// for i := 0; i < len(arr); i++ {
	// 	for j := 1; j < aim+1; j++ {
	// 		dp[i][j] = math.MaxInt32
	// 	}
	// }
	for index := len(arr) - 1; index >= 0; index-- {
		// 从第1列开始，第0列已经初始化
		for rest := 1; rest <= aim; rest++ {
			p1 := dp[index+1][rest]
			// rest-arr[index] 有可能会溢出，所以初始p2=-1
			p2 := -1
			if rest-arr[index] >= 0 {
				p2 = 1 + dp[index+1][rest-arr[index]]
			}

			if p1 == -1 && p2 == -1 {
				dp[index][rest] = -1
			} else {
				if p1 == -1 {
					dp[index][rest] = p2
				} else if p2 == -1 {
					dp[index][rest] = p1
				} else {
					dp[index][rest] = Min(p1, p2)
				}
			}
		}
	}

	return dp[0][aim]

}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

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
	lenght, max := 5, 10
	testime := 1
	for i := 0; i < testime; i++ {
		arr := generateRandomArry(lenght, max)
		aim := int(rand.Float32()*float32(max)*3) + max
		// fmt.Println(CoinMin3(arr, aim))
		if CoinMin2(arr, aim) != CoinMin3(arr, aim) {
			// fmt.Println(CoinMin2(arr, aim), CoinMin3(arr, aim))
			fmt.Println("xxx")
			break
		}
	}
}
