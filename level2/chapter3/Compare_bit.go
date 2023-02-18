package chapter3

/*
给定两个有符号32位整数a和b，返回a和b中较大的
tips:不用做任何比较判断
*/

// 反转
// 1 -> 0 , 0 -> 1
func filp(n int) int {
	return n ^ 1
}

// n 为非负数，返回 1
// n 为负数，返回 0
func sign(n int) int {
	return filp(n >> 31 & 1)
}

// a-b 可能会溢出，所以出现getMax2
func getMax(a, b int) int {
	c := a - b
	scA := sign(c)   // a-b 为负，scA为0; a-b非负，scA为1
	scB := filp(scA) // scA为0，scB为1; scA为1 ，scB为0
	// scA 和 scB 互斥
	return a*scA + b*scB
}

func getMax2(a, b int) int {
	c := a - b
	// 记录每个值的正负情况
	sa := sign(a)
	sb := sign(b)
	sc := sign(c)
	// a 和 b 符号不同，为1，相同为0
	difSab := sa ^ sb
	sameSab := filp(difSab) // 与difSab 互斥
	// 什么情况返回a，1) a和b符号不同,a >0  2) a和b符号相同，a-b >= 0
	returnA := difSab*sa + sameSab*sc
	returnB := filp(returnA)
	return a*returnA + b*returnB
}
