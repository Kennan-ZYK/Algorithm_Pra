package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func levelmaxNode(head *Node) {
	// 创建队列并存入根结点
	Queue := []*Node{head}
	// levelMap 各个结点的层数
	levelMap := map[*Node]int{}
	levelMap[head] = 1
	// 当前层数
	curlevel := 1
	// 当前层结点个数
	cntNode := 0
	// 最大结点个数
	max := -1
	for len(Queue) > 0 {
		// 弹出
		cur := Queue[0]
		Queue = Queue[1:]
		// 记录弹出cur的层数
		curNodelevel := levelMap[cur]
		// 判断结点是否属于当前层，属于curNode++,否则当前计算结束，进入下一层
		if curNodelevel == curlevel {
			cntNode++
		} else {
			// 结算上一层结点个数
			max = Max(max, cntNode)
			// 进入下一层
			curlevel++
			// 更新结点个数
			cntNode = 1
		}
		// 如果有左孩子，加入队列并记录所在层数
		if cur.Left != nil {
			levelMap[cur.Left] = curlevel + 1
			Queue = append(Queue, cur.Left)
		}
		// 如果有右孩子，加入队列并记录所在层数
		if cur.Right != nil {
			levelMap[cur.Right] = curlevel + 1
			Queue = append(Queue, cur.Right)
		}

	}
	// 结束循环后，结算最后一层
	max = Max(max, cntNode)
	return
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Print()
}
