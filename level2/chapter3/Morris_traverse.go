package main

import (
	"fmt"
)

/*
Morriss 先中后遍历
*/

type Node struct {
	val   int
	left  *Node
	right *Node
}

// 思路：只有一次直接打印，有两次的，遇见第一次打印
func MorrisPre(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		// 有左子树
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			// mostRight 变成左子树上最右结点
			if mostRight.right == nil { // 第一次来到cur
				fmt.Printf("%v", cur.val)
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // mostRight.right == cur
				mostRight.right = nil
			}

		} else { // 没有左子树
			fmt.Printf("%v", cur.val)
		}
		cur = cur.right

	}
}

// 思路：只有一次直接打印，有两次的，遇见第二次打印
func MorrisIn(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		// 有左子树
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			// mostRight 变成左子树上最右结点
			if mostRight.right == nil { // 第一次来到cur
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // mostRight.right == cur
				mostRight.right = nil
			}

		}
		// 打印
		fmt.Printf("%v", cur.val)
		cur = cur.right

	}
}

// 思路：只出现一次或者第一次遇见的结点都跳过，第二次遇见，逆序打印此结点左树右边界，遍历结束后，逆序打印整棵树的右边界
func MorrisPos(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		// 有左子树
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			// mostRight 变成左子树上最右结点
			if mostRight.right == nil { // 第一次来到cur
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // mostRight.right == cur
				mostRight.right = nil
				// 逆序打印此结点左树右边界
				printEdge(cur.left)
			}
		}
		cur = cur.right
	}
	// 打印整棵树的右边界
	printEdge(head)
}

func printEdge(x *Node) {
	// 逆序
	tail := reverseEdge(x)
	cur := tail
	for cur != nil {
		fmt.Printf("%v", cur.val)
		cur = cur.right
	}
	// 复原
	reverseEdge(tail)
}

func reverseEdge(from *Node) *Node {
	var pre *Node = nil
	var next *Node = nil
	for from != nil {
		next = from.right
		from.right = pre
		pre = from
		from = next
	}
	return pre
}

func main() {
	head := &Node{val: 1}
	head.left = &Node{val: 2}
	head.right = &Node{val: 3}

	head.left.left = &Node{val: 4}
	head.left.right = &Node{val: 5}
	head.right.left = &Node{val: 6}
	head.right.right = &Node{val: 7}

	head.left.left.right = &Node{val: 8}
	MorrisPre(head)
	fmt.Println()
	MorrisIn(head)
	fmt.Println()
	MorrisPos(head)

}
