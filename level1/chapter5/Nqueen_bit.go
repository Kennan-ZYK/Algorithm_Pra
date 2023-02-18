package chapter5

/*
和常规的N皇后区别就是，常数时间的效率会加快。时间复杂度还是O(n^n)

*/

// 最多32皇后，大于32 就把int类型改成long
func NQueen(n int) int {
	if n < 1 || n > 32 {
		return 0
	}
	var limit int
	if n == 32 {
		// 计算机中-1二进制是全1 ，符号位也是1
		limit = -1
	}
	limit = 1<<n - 1
	return process(limit, 0, 0, 0)
}

// colLim 列的限制，1的位置不能放皇后，0的位置可以
// leftDialim 左斜线的限制，1的位置不能放皇后，0的位置可以
// rightDiaLim 右斜线的限制，1的位置不能放皇后，0的位置可以
func process(limit, colLim, leftLim, rightLim int) int {
	// 表示所有位置都有皇后，完成1次
	if colLim == limit {
		return 1
	}
	// (colLim | leftLim | RightLim) 总限制。 ^0 = -1
	// 为什么要取反？因为我们定义的是 1 不能放，0能放
	// limit 取与 是为了去除符号位和截取我们要的n位(n=8,只需用到8位,其他标0)
	pos := limit & (^(colLim | leftLim | rightLim))
	lastRightOne := 0
	res := 0
	// 只看列不看行，一行过了就算
	for pos != 0 {
		// 取最右的1
		lastRightOne = pos & (^pos + 1)
		pos = pos - lastRightOne
		res += process(limit, lastRightOne|colLim, (lastRightOne|leftLim)<<1, (lastRightOne|rightLim)>>1)
	}
	return res
}
