package main

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func levelTreecnt(head *Node) int {
	if head == nil {
		return 0
	}
	Treemap := map[*Node]int{head: 1}
	Queue := []*Node{head}
	cnt := 0
	// 当前层
	curlevel := 1
	// 当前结点的层
	curNodelevel := 1
	max := math.MinInt64
	for len(Queue) > 0 {
		cur := Queue[0]
		Queue = Queue[1:]
		if Treemap[cur] == curlevel {
			cnt++
		} else {
			max = Max(max, cnt)
			cnt = 0
			// 下一层
			curlevel++
		}

		if cur.Left != nil {
			Queue = append(Queue, cur.Left)
			Treemap[cur.Left] = curlevel + 1
		}

		if cur.Right != nil {
			Queue = append(Queue, cur.Right)
			Treemap[cur.Right] = curlevel + 1
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
