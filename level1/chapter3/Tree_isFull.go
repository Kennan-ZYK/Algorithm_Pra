package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 记录信息
type Info struct {
	height int
	nodes  int
}

// 满二叉树 结点个数==2^(深度-1)
func isFull(head *Node) bool {
	if head == nil {
		return true
	}
	var data Info = process(head)
	// 1 << x == 2^x
	return data.nodes == (1<<data.height - 1)
}

func process(head *Node) Info {
	if head == nil {
		return Info{0, 0}
	}

	leftData := process(head.Left)
	rightData := process(head.Right)
	height := Max(leftData.height, rightData.height) + 1
	nodes := leftData.nodes + rightData.nodes + 1
	return Info{height, nodes}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	head := &Node{Val: 5}
	fmt.Printf("%T", head)
	head.Left = &Node{Val: 3}
	head.Right = &Node{Val: 7}
	head.Left.Left = &Node{Val: 2}
	head.Left.Right = &Node{Val: 4}
	head.Right.Left = &Node{Val: 6}
	head.Right.Right = &Node{Val: 8}
	ret := isFull(head)
	fmt.Println("是否是满二叉树：", ret)
}
