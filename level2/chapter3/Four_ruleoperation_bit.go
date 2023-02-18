package chapter3

// 四则运算
/*
给定两个有符号32位整数a和b，不能使用算术运算符，分别实现a和b的加、减、乘、除运算
要求
如果给定a、b执行加减乘除的运算结果就会导致数据的溢出，那么你实现的函数不必对此负责，除此之外请保证计算过程不发生溢出
*/

// ^异或 ，无进位相加，& << 1 进位相加
// 结束： & << 1 为 全 0
/*
 01101
 00111
 -------
 01010  ^ 异或
 00101  & 与
 01010  & << 1

 继续 把 异或 和 与左移一位，继续重复操作，直到 &为0
*/

// 参入参数溢出，不管
func Add(a, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
	return sum
}

// 减法
func neNum(n int) int {
	return Add(^n, 1)
}

func Minus(a, b int) int {
	return Add(a, neNum(b))
}

// 乘法
/*
	111     a
	101     b
   ------
	111
   000      相当于  a<<1 ,b>>1, 1110*0010,让第2位过去
  111
-----------
  11111
*/
func Multi(a, b int) int {
	c := uint64(b)
	res := 0
	for c != 0 {
		if c&1 != 0 { // 操作位 为1 就是1， 101 为1 ，100为0
			res = Add(res, a)
		}
		a = a << 1
		c = c >> 1
	}
	return res
}
