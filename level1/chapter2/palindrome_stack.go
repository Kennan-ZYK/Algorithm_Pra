package main

import "fmt"

type Stack struct {
	MaxSize int
	Arr     []int
	Top     int
}

func (this *Stack) Push(val int) {
	if this.Top == this.MaxSize-1 {
		fmt.Println("栈满")
	}
	this.Top++
	this.Arr = append(this.Arr, val)
}

func (this *Stack) POP() (val int) {
	if this.Top == -1 {
		fmt.Println("栈空")
	}
	val = this.Arr[this.Top]
	this.Top--
	return val
}

func (this *Stack) Display() {
	if this.Top == -1 {
		fmt.Println("栈空")
	}
	for i := this.Top; i > -1; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.Arr[i])
	}

}

func (this *Stack) ispalindrome(nums []int) bool {
	if this.Top == -1 {
		fmt.Println("栈空无法判断是否为回文数")
	}

	for i, j := this.Top, 0; i > -1; i, j = i-1, j+1 {
		ans := this.POP()
		if nums[j] != ans {
			return false
		}
	}

	return true
}

//采用n/2空间完成
// func (this *Stack) ispalindrome_new(nums []int) bool {
// 	if this.Top == -1 {
// 		fmt.Println("栈空无法判断是否为回文数")
// 	}
// }

func main() {
	stack := &Stack{
		MaxSize: 10,
		Top:     -1,
	}
	nums := []int{1, 2, 3, 2, 1}
	for i := 0; i < len(nums); i++ {
		stack.Push(nums[i])
	}
	// stack.Display()
	fmt.Println(stack.ispalindrome(nums))
}
