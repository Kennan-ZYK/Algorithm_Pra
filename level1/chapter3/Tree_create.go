package main

import (
	"fmt"
)

type btNode struct {
	Val    interface{}
	Lchild *btNode
	Rchild *btNode
}

type biTree struct {
	root *btNode
}

func NewTree() *biTree {
	return &biTree{}
}

func newTreeNode(data interface{}) *btNode {
	return &btNode{
		Val:    data,
		Lchild: nil,
		Rchild: nil,
	}
}
func (bt *biTree) addNode(data interface{}) {
	cur := bt.root
	if cur == nil { //空树则创建一个根结点
		bt.root = newTreeNode(data)
		return
	}
	Queue := []*btNode{cur}
	// 更新头结点 0->1，依次类推
	/*    	0
		1		2
	  3   4   5	  6
	*/
	for len(Queue) > 0 {
		cur := Queue[0]
		// fmt.Println(cur)
		Queue = Queue[1:]
		// fmt.Println(Queue)
		if cur.Lchild != nil {
			Queue = append(Queue, cur.Lchild)
		} else {
			cur.Lchild = newTreeNode(data)
			return
		}
		if cur.Rchild != nil {
			Queue = append(Queue, cur.Rchild)
		} else {
			cur.Rchild = newTreeNode(data)
			break
		}
	}
}

func preOrderUnRecur(head *btNode) {
	// stack := []*btNode{}
	// for head != nil || len(stack) > 0 {
	// 	for head != nil {
	// 		fmt.Printf("%v ", head.Val)
	// 		stack = append(stack, head)
	// 		head = head.Lchild
	// 	}
	// 	head = stack[len(stack)-1].Rchild
	// 	stack = stack[:len(stack)-1]
	// }
	// 左神写法 //
	if head != nil {
		stack := []*btNode{}
		stack = append(stack, head)
		for len(stack) > 0 {
			head = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%v ", head.Val)
			if head.Rchild != nil {
				stack = append(stack, head.Rchild)
			}
			if head.Lchild != nil {
				stack = append(stack, head.Lchild)
			}
		}
	}
	fmt.Println()
}

// 计算一层中最多结点数
func levelmaxNode(head *btNode) int {
	Queue := []*btNode{head}
	// levelMap 各个结点的层数
	levelMap := map[*btNode]int{}
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
		if cur.Lchild != nil {
			levelMap[cur.Lchild] = curlevel + 1
			Queue = append(Queue, cur.Lchild)
		}
		// 如果有右孩子，加入队列并记录所在层数
		if cur.Rchild != nil {
			levelMap[cur.Rchild] = curlevel + 1
			Queue = append(Queue, cur.Rchild)
		}

	}
	// 结束循环后，结算最后一层
	max = Max(max, cntNode)
	return max
}

func levelNode(head *btNode) int {
	var curend *btNode = head
	var nextend *btNode = nil
	Queue := []*btNode{head}
	cnt := 0
	max := -1
	for len(Queue) > 0 {
		cur := Queue[0]
		Queue = Queue[1:]
		cnt++
		if nextend == nil {
			nextend = cur
		}
		if cur.Lchild != nil {
			Queue = append(Queue, cur.Lchild)
			nextend = cur.Lchild
		}
		if cur.Rchild != nil {
			Queue = append(Queue, cur.Rchild)
			nextend = cur.Rchild
		}
		if curend == cur {
			max = Max(max, cnt)
			curend = nextend
			// nextend = nil
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

func main() {

	list := []interface{}{0, 1, 2, 3, 4, 5}
	// node := []*btNode{}
	tree := NewTree()
	for _, Val := range list {
		// node = append(node, newTreeNode(Val))
		tree.addNode(Val)
	}

	// tree.root = node[0]
	// tree.root.Lchild = node[1]
	// tree.root.Rchild = node[2]
	// tree.root.Lchild.Lchild = node[3]
	// tree.root.Lchild.Rchild = node[4]
	// tree.root.Rchild.Lchild = node[5]
	// preOrderUnRecur(tree.root)
	// inOrderUnRecur(tree.root)
	// fmt.Print(levelmaxNode(tree.root))
	fmt.Print(levelNode(tree.root))
}
