package chapter4

/*
N,S,E,K
一个机器人 在 有1~N个位置的路行走，可以向左向右移动，每次能移动一步，在左边界时，只能向右移动，
在右边界时，只能向左移动， 起始点为S, 目标位置为E, 能走多少步 K
现在要求 机器人在 S点出发，在能走K步的情况下，有几种走的路线到E
*/

// 方法1 ：暴力递归 O(2^K)
func walkWay1(N, S, E, K int) int {
	if N < 2 || K < 1 || S < 1 || E > N || S > N || E < 1 {
		return 0
	}
	return f1(N, E, K, S)
}

// 一共有 1~N位置，
// 目标位置 E
// rest 剩余步数
// cur 当前位置
func f1(N, E, rest, cur int) int {
	// base case
	if rest == 0 {
		if cur == E {
			return 1
		} else {
			return 0
		}
	}
	// 左边界
	if cur == 1 {
		return f1(N, E, rest-1, 2)
	}
	// 右边界
	if cur == N {
		return f1(N, E, rest-1, N-1)
	}

	return f1(N, E, rest-1, cur-1) + f1(N, E, rest-1, cur+1)
}

// 方法2 : 记忆搜索 O(KN)
/*
因为方法1 会有重复计算，
如12345， S = 2, K = 4
		2
	1		3
	2*	  2*   4
...
2* 标星的2，他们partion的结果一样而且rest =2 ，重复的都是1 3
所以记 dp[rest][cur] = dp[2][2] = 下面计算过的情况和
*/

func walkWay2(N, S, E, K int) int {
	if N < 2 || K < 1 || S < 1 || E > N || S > N || E < 1 {
		return 0
	}
	// 开辟一个缓存存储出现过的情况
	var dp [][]int
	dp = make([][]int, K+1)
	// 初始化
	for i := range dp {
		// 0 位置不用，题意说1~N，所以+1
		dp[i] = make([]int, N+1)
	}

	for i := 0; i < K; i++ {
		for j := 0; j < N; j++ {
			dp[i][j] = -1
		}
	}
	return f2(N, E, K, S, dp)
}

func f2(N, E, rest, cur int, dp [][]int) int {
	// 如果记录过，直接返回
	if dp[rest][cur] != -1 {
		return dp[rest][cur]
	}
	if rest == 0 {
		if cur == E {
			dp[rest][cur] = 1
		} else {
			dp[rest][cur] = 0
		}
		// return dp[rest][cur]
	} else if cur == 1 {
		dp[rest][cur] = f2(N, E, rest-1, 2, dp)
		// return dp[rest][cur]
	} else if cur == E {
		dp[rest][cur] = f2(N, E, rest-1, N-1, dp)
		// return dp[rest][cur]
	} else {
		dp[rest][cur] = f2(N, E, rest-1, cur-1, dp) + f2(N, E, rest-1, cur+1, dp)
	}

	return dp[rest][cur]
}

// 方法3：动态规划 O(KN)
func dpWay(N, S, E, K int) int {
	if N < 2 || K < 1 || S < 1 || E > N || S > N || E < 1 {
		return 0
	}
	var dp [][]int
	dp = make([][]int, K+1)

	// 初始化
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	// 结束位置
	dp[0][E] = 1

	//
	for rest := 1; rest <= K; rest++ {
		for cur := 1; cur <= N; cur++ {
			if cur == 1 {
				dp[rest][cur] = dp[rest-1][2]
			} else if cur == N {
				dp[rest][cur] = dp[rest-1][N-1]
			} else {
				dp[rest][cur] = dp[rest-1][cur-1] + dp[rest-1][cur+1]
			}
		}
	}
	// 返回开始点，为什么呢? 因为我们是反推得出结果的，可以看robot图片了解
	return dp[S][K]
}
