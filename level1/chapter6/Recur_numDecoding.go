package chapter6

/*
题目:规定1和A对应，2和B对应，3和C对应...
有一个数字字符串比如"111",可以转化成"AAA"、"KA"、"AK"。
给定一个只有数字字符组成的字符串str，返回有多少种转化结果
*/

func numDecoding(str string) int {
	if len(str) == 0 {
		return 0
	}

	res := process(str, 0, 0)
	return res

}

// i是从第0位置开始
func process(str string, i, res int) int {
	if len(str) == i {
		return 1
	}
	if str[i] == '0' {
		return 0
	}

	if str[i] == '1' {
		process(str, i+1, res)
		if i+1 < len(str) {
			res += process(str, i+2, res)
		}
		return res
	}
	if str[i] == '2' {
		process(str, i+1, res)
		if i+1 < len(str) && (str[i+1] >= '0' && str[i+1] < '6') {
			process(str, i+2, res)
		}
		return res
	}

	return process(str, i+1, res)
}
