package main

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

//L1 ，L2都是头结点
func Linkcommon(L1 *LinkNode, L2 *LinkNode) []int {
	l1 := L1.Next
	l2 := L2.Next
	ans := []int{}
	for l1 != nil || l2 != nil {
		if l1.Val < l2.Val {
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			l2 = l2.Next
		} else {
			ans = append(ans, l1.Val)
		}
	}
	return ans
}

func main() {
	fmt.Print("Ss")
}
