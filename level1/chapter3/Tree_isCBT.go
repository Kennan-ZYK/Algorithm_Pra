package chapter3

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 判断条件：不是完全二叉树
// 1. 任一结点有右无左， false
// 2. 在1条件的前提，遇到第一个左子不全，后续皆为子结点

func isCBT(head *Node) bool {
	if head == nil {
		return true
	}
	// 条件2 开关
	leaf := false
	Queue := []*Node{head}
	for len(Queue) > 0 {
		head = Queue[0]
		Queue = Queue[1:]
		l := head.Left
		r := head.Right
		// 有右无左 或者 开关触发后，剩下不为叶子结点
		if l == nil && r != nil || leaf && (l != nil || r != nil) {
			return false
		}
		if l != nil {
			Queue = append(Queue, l)
		}
		if r != nil {
			Queue = append(Queue, r)
		}
		// 左有右无，剩下必须为叶子结点 l==nil(可省) || r==nil
		if r == nil {
			leaf = true
		}
	}
	return true
}

// 方法2：bfs

func isCompleteTree(head *Node) bool {
	//标记层序遍历的过程中是否有遇到空节点
	empty := false
	//辅助层序遍历的队列
	queue := []*Node{head}
	for len(queue) > 0 {
		//取一个节点
		node := queue[0]
		//出队列
		queue = queue[1:]
		if node == nil {
			//遇到空节点，把标记改成true
			empty = true
		} else {
			//此时是遍历非空节点，在非空节点之前出现了空节点，就说明并不是满二叉树
			if empty == true {
				return false
			}
			//添加左右子节点，无论是否为空
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return true
}
