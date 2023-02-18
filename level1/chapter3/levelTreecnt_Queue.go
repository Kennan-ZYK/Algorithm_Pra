package chapter3

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func levelmaxNode(head *Node) int {
	// 创建保存当前层最后一个结点
	var curend *Node = head
	// 创建保存下层最后一个结点
	var nextend *Node = nil
	// 创建队列并存入根结点
	Queue := []*Node{head}
	// 计数器
	cnt := 0
	// 最大结点数
	max := -1
	for len(Queue) > 0 {
		// 弹出
		cur := Queue[0]
		Queue = Queue[1:]
		// 结点加一
		cnt++
		// 初始赋值，若定义时赋值了head，可以省略
		if nextend == nil {
			nextend = cur
		}
		// 左子树不为空，加入队列，更新nextend
		if cur.Left != nil {
			Queue = append(Queue, cur.Left)
			nextend = cur.Left
		}
		// 右子树不为空，加入队列，更新nextend
		if cur.Right != nil {
			Queue = append(Queue, cur.Right)
			nextend = cur.Right
		}
		// 一层结束，统计最大结点max，把下一层的最后结点更新成当前层最后结点
		if curend == cur {
			max = Max(max, cnt)
			curend = nextend
			nextend = nil
			// 重新计数
			cnt = 0
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
