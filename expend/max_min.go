package main

import "fmt"

func main() {

	const UINT_MIN uint = 0
	// 其最大值的二进制表示的所有位都为1，那么，
	const UINT_MAX = ^uint(0)
	// 有符号整型int
	// fmt.Printf("%+V", UINT_MAX)
	// 根据补码，其最大值二进制表示，首位0，其余1，那么，
	const INT_MAX = int(^uint(0) >> 1)
	fmt.Printf("%b", INT_MAX)
	// 根据补码，其最小值二进制表示，首位1，其余0，那么，
	const INT_MIN = ^INT_MAX

	// 还有其他数据类型与上述同理

	// 还可通过进制转换查看转换流程

	// 十进制转二进制 %d -> %b（其余同理）
	// fmt.Printf("%d的十进制为%b", num01, num01)
}
