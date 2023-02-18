package main

import "fmt"

/*
二叉树结点的最大距离问题
从二叉树的结点a出发，可以向上或向下走，但沿途的结点只能经过一次，到达结点b，经过的结点个数叫做x到y的距离，
那么二叉树任何两个结点之间都有距离，求整棵树的最大距离。
情况1：x不参与
左的最大距离 或者 右的最大距离
如：最大距离是 a o 1 o b
			x
		1		o
	o 		o
a  			   b
情况2： x参与
左高度 + 1 + 右高度
				x
    	1				2
	3       4		5 		6
a 								b
最大距离 a->3->1->x->2->6->b
*/

type Node struct {
	val   int
	left  *Node
	right *Node
}

type info struct {
	Maxdepth int
	height   int
}

func MaxDistance(x *Node) int {
	res := process(x)
	return res.Maxdepth
}

func process(head *Node) info {
	if head == nil {
		return info{0, 0}
	}
	leftInfo := process(head.left)
	rightInfo := process(head.right)

	// 3种情况
	p1 := leftInfo.Maxdepth
	p2 := rightInfo.Maxdepth
	p3 := leftInfo.height + 1 + rightInfo.height

	MaxDis := Max(p3, Max(p1, p2))
	// +1 是因为还要加上x结点
	height := Max(leftInfo.height, rightInfo.height) + 1
	return info{MaxDis, height}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	head := &Node{val: 1}
	head.left = &Node{val: 2}
	head.right = &Node{val: 3}
	head.left.left = &Node{val: 4}
	head.left.right = &Node{val: 5}
	head.left.left.left = &Node{val: 5}
	head.left.left.left.left = &Node{val: 5}
	head.left.right.left = &Node{val: 5}
	// head.right.left = &Node{val: 6}
	// head.right.right = &Node{val: 7}
	fmt.Println(MaxDistance(head))
}
