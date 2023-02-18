package main

import "fmt"

/*
在象棋盘中，起点为（0，0）,马刚好走step步到(x,y)，问有多少种方法能到？
象棋盘是 9*10
*/

// 思路：逆推过程，从终点退回起点
func walkway(x, y, step int) int {
	if x < 0 || x > 8 || y < 0 || y > 9 {
		return 0
	}
	if step == 0 {
		if x == 0 && y == 0 {
			return 1
		} else {
			return 0
		}
	}

	return walkway(x-1, y+2, step-1) +
		walkway(x+1, y+2, step-1) +
		walkway(x+2, y+1, step-1) +
		walkway(x+2, y-1, step-1) +
		walkway(x+1, y-2, step-1) +
		walkway(x-1, y-2, step-1) +
		walkway(x-2, y-1, step-1) +
		walkway(x-2, y+1, step-1)

}

// 动态规划
// 这题有三个自变量，所以是一个体
func dpway(x, y, step int) int {
	if x < 0 || x > 8 || y < 0 || y > 9 || step < 0 {
		return 0
	}

	// 长 0~8
	dp := make([][][]int, 9)
	// 宽
	for i := range dp {
		dp[i] = make([][]int, 10)
	}
	// 高
	for i := 0; i < 9; i++ {
		for j := 0; j < 10; j++ {
			dp[i][j] = make([]int, step+1)
		}
	}

	// 起点
	dp[0][0][0] = 1
	for h := 1; h <= step; h++ {
		for r := 0; r < 9; r++ {
			for c := 0; c < 10; c++ {
				dp[r][c][h] += getVaild(dp, r-1, c+2, h-1)
				dp[r][c][h] += getVaild(dp, r+1, c+2, h-1)
				dp[r][c][h] += getVaild(dp, r+2, c+1, h-1)
				dp[r][c][h] += getVaild(dp, r+2, c-1, h-1)
				dp[r][c][h] += getVaild(dp, r+1, c-2, h-1)
				dp[r][c][h] += getVaild(dp, r-1, c-2, h-1)
				dp[r][c][h] += getVaild(dp, r-2, c-1, h-1)
				dp[r][c][h] += getVaild(dp, r-2, c+1, h-1)
			}
		}
	}

	return dp[x][y][step]
}

func getVaild(dp [][][]int, row, col, step int) int {
	if row < 0 || row > 8 || col < 0 || col > 9 {
		return 0
	}
	return dp[row][col][step]
}

func main() {
	x := 7
	y := 7
	step := 10
	// dpway(x, y, step)
	fmt.Printf("walkway(x, y, step): %v\n", walkway(x, y, step))
	fmt.Printf("walkway(x, y, step): %v\n", dpway(x, y, step))
}
