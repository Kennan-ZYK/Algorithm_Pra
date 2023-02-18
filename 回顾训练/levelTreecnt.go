package main

import "math"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func levelTreecnt(head *Node) int {
	if head == nil {
		return 0
	}
	Queue := []*Node{head}
	endNode := head
	cnt := 0
	curend := head
	max := math.MinInt64
	for len(Queue) > 0 {
		// 弹出结点
		cur := Queue[0]
		Queue = Queue[1:]
		cnt++
		if cur.left != nil {
			Queue = append(Queue, cur.left)
			endNode = cur.left
		}

		if cur.right != nil {
			Queue = append(Queue, cur.right)
			endNode = cur.right
		}
		// 一层结束
		if endNode == cur {
			max = Max(max, cnt)
			cnt = 0
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
