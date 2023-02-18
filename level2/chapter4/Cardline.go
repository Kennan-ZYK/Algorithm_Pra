package main

import "fmt"

/*
	在level1文件夹的chapte6的dealcard是暴力写法，现在用记忆搜索法和动态规划
*/

func win(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return Max(f(arr, 0, len(arr)-1), s(arr, 0, len(arr)-1))
}

func f(arr []int, L, R int) int {
	// 先手拿
	if L == R {
		return arr[L]
	}
	return Max(arr[L]+s(arr, L+1, R), arr[R]+s(arr, L, R-1))
}

func s(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	return Min(f(arr, L+1, R), f(arr, L, R-1))
}

// 记忆化搜索
func win2(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	N := len(arr)
	fdp := make([][]int, N)
	sdp := make([][]int, N)
	// 初始化
	for i := range fdp {
		fdp[i] = make([]int, N)
		sdp[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fdp[i][j] = -1
			sdp[i][j] = -1
		}
	}
	return Max(f2(arr, 0, N-1, fdp, sdp), s2(arr, 0, N-1, fdp, sdp))
}

func f2(arr []int, left, right int, fdp, sdp [][]int) int {
	if fdp[left][right] != -1 {
		return fdp[left][right]
	}
	ans := 0
	if left == right {
		return arr[left]
	} else {
		ans = Max(arr[left]+s2(arr, left+1, right, fdp, sdp), arr[right]+s2(arr, left, right-1, fdp, sdp))
	}
	fdp[left][right] = ans
	return ans
}

func s2(arr []int, left, right int, fdp, sdp [][]int) int {
	if sdp[left][right] != -1 {
		return sdp[left][right]
	}
	ans := 0
	if left != right {
		ans = Min(f2(arr, left+1, right, fdp, sdp), f2(arr, left, right-1, fdp, sdp))
	}
	sdp[left][right] = ans
	return ans

}

// 动态规划
func win3(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	N := len(arr)
	fdp := make([][]int, N)
	sdp := make([][]int, N)
	// 初始化
	for i := range fdp {
		fdp[i] = make([]int, N)
		sdp[i] = make([]int, N)
	}
	// 根据 i == j return arr[i]，可得
	for i := range fdp {
		fdp[i][i] = arr[i]
	}
	for stratCol := 1; stratCol < N; stratCol++ {
		L := 0
		R := stratCol
		for R < N {
			fdp[L][R] = Max(arr[L]+sdp[L+1][R], arr[R]+sdp[L][R-1])
			sdp[L][R] = Min(fdp[L+1][R], fdp[L][R-1])
			L++
			R++
		}
		fmt.Println("第", stratCol, "次")
		fmt.Println(fdp, "\n", sdp)
	}

	return Max(fdp[0][N-1], sdp[0][N-1])
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	arr := []int{1, 4, 5, 4}
	// 4+4
	// 5+1
	fmt.Print(win3(arr))
	// if win(arr) == win2(arr) && win2(arr) == win3(arr) {
	// 	fmt.Print(12)
	// } else {
	// 	fmt.Print("xx")
	// }
}
