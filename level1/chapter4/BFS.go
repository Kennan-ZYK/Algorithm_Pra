package chapter4

import "fmt"

type Node struct {
	Val   int
	nexts []*Node
}

// 宽度优先遍历
/*
1、利用队列实现
2、从源结点开始把结点按照宽度放入队列，然后弹出
3、每弹出一个点，把该结点下一个没有进过队列的邻接点放入栈
4、直到队列变空
*/
// 无向图
func BFS(node *Node) {
	if node == nil {
		return
	}
	Queue := []*Node{node}
	// 记录结点，避免重复
	set := map[*Node]bool{node: true}
	for len(Queue) > 0 {
		cur := Queue[0]
		Queue = Queue[1:]
		fmt.Printf("%v", cur.Val)
		for i := 0; i < len(cur.nexts); i++ {
			if !set[cur.nexts[i]] {
				Queue = append(Queue, cur.nexts[i])
				set[cur.nexts[i]] = true
			}
		}
	}
}
