package chapter1

/*
给定一个由 0 和 1 组成的非空二维数组 grid ，用来表示海洋岛屿地图。
一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
你可以假设 grid 的四个边缘都被 0（代表水）包围着。
找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 0
*/

func maxAreaOfIsland(grid [][]int) int {
	L, C := len(grid), len(grid[0])
	res := 0
	max := 0
	// ans = 0
	for i := 0; i < L; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 1 {
				res = infect(grid, i, j, L, C)
				max = Max(max, res)
			}
		}
	}
	return max
}

// ans变量可以不用，return 1+... 即可，因为只要进入 infect函数 就一定存在1个岛
// O(L*C)
func infect(grid [][]int, i, j, L, C int) (ans int) {
	// 边界实际下标要-1
	if i < 0 || i >= L || j < 0 || j >= C || grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 2
	ans++

	return ans + infect(grid, i+1, j, L, C) + infect(grid, i-1, j, L, C) + infect(grid, i, j+1, L, C) + infect(grid, i, j-1, L, C)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
