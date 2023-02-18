package chapter2

func dailyTemperatures(temperatures []int) []int {
	// 单调栈记录下标
	stack := []int{}
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	// 剩下的都是后面没有大于自己的
	for len(stack) != 0 {
		ans[stack[len(stack)-1]] = 0
		stack = stack[1:]
	}
	return ans
}
