package chapter6

/*题目：逆序栈 不用额外的数据结构，只要用递归*/

// 思路：利用递归思想，每一层递归都会保留数据
// tips：画图很好理解
func reverse(stack []int) {
	if len(stack) == 0 {
		return
	}
	res := f(stack)
	reverse(stack)
	Push(&stack, res)
}

// 把栈底元素取出，其他元素顺序不变
func f(stack []int) int {
	res := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	if len(stack) == 0 {
		return res
	} else {
		last := f(stack)
		// 压栈
		Push(&stack, res)
		return last
	}

}

func Push(stack *[]int, res int) {
	(*stack) = append((*stack), res)
}
