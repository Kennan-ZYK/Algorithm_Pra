package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Info struct {
	isBTS bool
	min   int
	max   int
}

func isBTS(head *Node) bool {
	if head == nil {
		return true
	}
	res := process(head)
	// fmt.Print(res)
	return res.isBTS
}

func process(head *Node) Info {
	//
	if head == nil {
		return Info{false, 0, 0}
	}
	leftData := process(head.Left)
	rightData := process(head.Right)
	min := head.Val
	max := head.Val
	// 左树不为空
	if !IsEmpty(leftData) {
		max = Max(max, leftData.max)
		min = Min(min, leftData.min)
	}
	// 右树不为空
	if !IsEmpty(rightData) {
		max = Max(max, rightData.max)
		min = Min(min, rightData.min)
	}

	isBTS := true
	// 左树最大值要小于当前值才是搜索树
	if !IsEmpty(leftData) && (!leftData.isBTS || leftData.max <= head.Val) {
		isBTS = false
	}
	// 右树最小值要大于当前值才是搜索树
	if !IsEmpty(rightData) && (!rightData.isBTS || rightData.min <= head.Val) {
		isBTS = false
	}
	return Info{isBTS, min, max}
}

func IsEmpty(a Info) bool {
	// 判断 a == nil
	return reflect.DeepEqual(a, Info{})
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	head := &Node{Val: 2}
	head.Left = &Node{Val: 7}
	head.Right = &Node{Val: 3}
	// head.Left.Left = &Node{Val: 2}
	// head.Left.Right = &Node{Val: 4}
	// head.Right.Left = &Node{Val: 6}
	// head.Right.Right = &Node{Val: 8}
	ret := isBTS(head)
	fmt.Println("是否是搜索二叉树：", ret)
}
