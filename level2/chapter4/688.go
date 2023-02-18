package chapter4

import (
	"fmt"
	"math"
)

/*
在一个 n x n 的国际象棋棋盘上，一个骑士从单元格 (row, column) 开始，并尝试进行 k 次移动。
行和列是 从 0 开始 的，所以左上单元格是 (0,0) ，右下单元格是 (n - 1, n - 1) 。
象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。
*/

// 暴力递归
// n = 8 k = 30 row = 6 column = 4
func knightProbability(n int, k int, row int, column int) float64 {
	ans := process(n, k, row, column)
	fmt.Println(ans)
	result := math.Round(float64(ans)/math.Pow(float64(8), float64(k))*1e5) / 1e5
	return result
}

func process(n, k, row, column int) int {
	if column < 0 || column >= n || row >= n || row < 0 {
		return 0
	}
	if k == 0 {
		return 1
	}
	live := process(n, k-1, row+1, column+2)
	live += process(n, k-1, row-1, column+2)
	live += process(n, k-1, row+2, column+1)
	live += process(n, k-1, row-2, column+1)
	live += process(n, k-1, row+2, column-1)
	live += process(n, k-1, row-2, column-1)
	live += process(n, k-1, row+1, column-2)
	live += process(n, k-1, row-1, column-2)
	return live
}

// 记忆搜索法
func knightProbability2(n, k, row, column int) float64 {
	// 初始化表
	dp := make([][][]float64, n)

	for i := 0; i < n; i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]float64, k+1)
		}
	}
	ans := process1(n, k, row, column, dp)
	result := math.Round(ans/math.Pow(float64(8), float64(k))*1e5) / 1e5
	return result

}

func process1(n, k, row, column int, dp [][][]float64) float64 {
	if row < 0 || row >= n || column < 0 || column >= n {
		return 0
	}
	if k == 0 {
		return 1
	}
	if dp[row][column][k] != 0 {
		return dp[row][column][k]
	}

	dp[row][column][k] += process1(n, k-1, row+1, column+2, dp)
	dp[row][column][k] += process1(n, k-1, row-1, column+2, dp)
	dp[row][column][k] += process1(n, k-1, row+2, column+1, dp)
	dp[row][column][k] += process1(n, k-1, row-2, column+1, dp)
	dp[row][column][k] += process1(n, k-1, row+2, column-1, dp)
	dp[row][column][k] += process1(n, k-1, row-2, column-1, dp)
	dp[row][column][k] += process1(n, k-1, row+1, column-2, dp)
	dp[row][column][k] += process1(n, k-1, row-1, column-2, dp)

	return dp[row][column][k]

}

// 动态规划

func knightProbability3(n, k, row, column int) float64 {
	// 初始化
	dp := make([][][]float64, n)

	for i := 0; i < n; i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]float64, k+1)
			for q := 0; q < k+1; q++ {
				dp[i][j][0] = 1
			}
		}
	}
	way := [][]int{{-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {-2, 1}, {-2, -1}, {2, 1}, {2, -1}}

	// 第0层已经初始化，从第一层开始
	for h := 1; h < k+1; h++ {
		for row := 0; row < n; row++ {
			for column := 0; column < n; column++ {
				for _, w := range way {
					x, y := row+w[0], column+w[1]
					// 判断是否有效
					if getVaild(n, x, y) {
						dp[row][column][h] += dp[x][y][h-1] / 8
					}
				}
			}
		}
	}

	return dp[row][column][k]

}

func getVaild(n, row, column int) bool {
	if row < 0 || row >= n || column < 0 || column >= n {
		return false
	}
	return true
}
