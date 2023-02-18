package main

import "math"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func Morris_isBTS(head *Node) bool {
	if head == nil {
		return true
	}
	cur := head
	var mostRight *Node = nil
	preVal := math.MinInt64
	for cur != nil {
		mostRight = cur.left
		// 如果有左子树，则找到其左子树最右结点
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else { //mostRight.right == cur
				mostRight.right = nil
			}
		}
		if preVal <= cur.val {
			return false
		}
		preVal = cur.val
		cur = cur.right
	}
	return true
}
