package expend

func test(R, L int) {
	// 写法1：当R和L很大时，mid会溢出
	mid := (R + L) / 2
	// 写法2
	mid = L + (R-L)/2
	// 写法3
	mid = L + (R-L)>>1
}
