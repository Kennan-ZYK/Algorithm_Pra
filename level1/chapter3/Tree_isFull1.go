package chapter3

import "fmt"

func main() {
	head := &TreeNode{Val: 5}
	fmt.Printf("%T", head)
	head.Left = &TreeNode{Val: 3}
	head.Right = &TreeNode{Val: 7}
	head.Left.Left = &TreeNode{Val: 2}
	head.Left.Right = &TreeNode{Val: 4}
	head.Right.Left = &TreeNode{Val: 6}
	// head.Right.Right = &TreeNode{Val: 8}
	ret := IsFBT(head)
	fmt.Println("是否是满二叉树：", ret)
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Info struct {
	IsFull bool
	Cnt    int
}

func IsFBT(head *TreeNode) bool {
	return process(head).IsFull
}

func process(head *TreeNode) *Info {
	if head == nil {
		return &Info{IsFull: true}
	}

	leftInfo := process(head.Left)
	//左不满
	if !leftInfo.IsFull {
		return leftInfo
	}

	rightInfo := process(head.Right)
	//右不满
	if !rightInfo.IsFull {
		return rightInfo
	}

	//左右不等
	if leftInfo.Cnt != rightInfo.Cnt {
		return new(Info)
	}

	//通过所有考验
	return &Info{IsFull: true, Cnt: leftInfo.Cnt + rightInfo.Cnt + 1}

}
