package chapter4

import "fmt"

type Node struct {
	Val   int
	nexts []*Node
}

// 广度优先遍历
/*
1、利用栈实现
2、从源结点开始把结点按照深度放入栈，然后弹出
3、每弹出一个点，把该结点下一个没有进过栈的邻接点放入栈
4、直到栈变空
*/
// 无向图
func DFS(node *Node) {
	if node == nil {
		return
	}
	Stack := []*Node{node}
	set := map[*Node]bool{node: true}
	for len(Stack) > 0 {
		cur := Stack[len(Stack)-1]
		Stack = Stack[:len(Stack)-1]
		fmt.Printf("%v", node.Val)
		for i := 0; i < len(node.nexts); i++ {
			if !set[node.nexts[i]] {
				Stack = append(Stack, cur)
				Stack = append(Stack, node.nexts[i])
				set[node.nexts[i]] = true
				fmt.Printf("%v", node.nexts[i].Val)
				break
			}
		}

	}

}
