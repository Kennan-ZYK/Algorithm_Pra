package main

import "math"

/*N皇后
在N*N的棋盘上要摆N个皇后，要求任何两个皇后不同行、不同列，也不在同一条斜线上。
给定一个整数n,返回n皇后的摆法有多少种。
n=1,返回1
n=2或3，2皇后和3皇后怎么摆都不行，返回0
n=8，返回92
*/

func NQueen(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return process(0, record, n)
}

// i 表示当前来到第几行
// record[i]  i行的皇后保存在第几列
// n 表示总共有几行
func process(i int, record []int, n int) int {
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ {
		// 当前i行的皇后，放在j列是否有效，会不会共斜线或者共列
		// 如果有效，则放入皇后
		if Valid(record, i, j) {
			record[i] = j
			res += process(i+1, record, n)
		}
	}
	return res
}

func Valid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		// record[k] == j 判断是否共列
		// k-i  ，i为当前行，k为之前的有效行。record[k] 之前行的列
		// 共斜线就是斜率为45度或者135度，abs(y1-y2 / x1-x2) == 1
		// record[k] - j 不会等于0 ，因为record[k]==j已经被相等情况排除了
		// 也可以写成math.Abs(float64(k-i)) == math.Abs(float64(record[k]-j))
		if record[k] == j || math.Abs(float64((k-i)/record[k]-j)) == 1 {
			//
			return false
		}
	}
	return true
}
