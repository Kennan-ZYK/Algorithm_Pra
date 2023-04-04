package chapter3

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//递归
func preOrderRecur(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("%v ", head.Val)
	preOrderRecur(head.Left)
	preOrderRecur(head.Right)
}

func posOrderRecur(head *Node) {
	if head == nil {
		return
	}
	posOrderRecur(head.Left)
	posOrderRecur(head.Right)
	fmt.Printf("%v ", head.Val)
}

func inOrderRecur(head *Node) {
	if head == nil {
		return
	}
	inOrderRecur(head.Left)
	fmt.Printf("%v ", head.Val)
	inOrderRecur(head.Right)
}

func inorderTraversal(root *Node) (res []int) {
	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

func leveOrder(head *Node) {
	if head == nil {
		return
	}
	// 先放入根结点
	Queue := []*Node{head}
	for len(Queue) > 0 {
		// 更新头结点, 队列先进先出
		head = Queue[0]
		Queue = Queue[1:]
		fmt.Printf("%v ", head.Val)
		if head.Left != nil {
			Queue = append(Queue, head.Left)
		}
		if head.Right != nil {
			Queue = append(Queue, head.Right)
		}
	}
}

// 非递归

/*前序：头左右
1、从栈中弹出一个节点cur
2、打印（处理）cur
3、先右再左（如果有）放入栈
4、重复
*/
func preOrderUnRecur(head *Node) {
	stack := []*Node{}
	for head != nil || len(stack) > 0 {
		for head != nil {
			fmt.Printf("%v ", head.Val)
			stack = append(stack, head)
			head = head.Left
		}
		head = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	// 左神写法 //
	// if head != nil {
	// 	stack := []*Node{}
	// 	stack = append(stack, head)
	// 	for len(stack) > 0 {
	// 		head = stack[len(stack)-1]
	// 		stack = stack[:len(stack)-1]
	// 		fmt.Printf("%v->", head)
	// 		if head.Right != nil {
	// 			stack = append(stack, head.Right)
	// 		}
	// 		if head.Left != nil {
	// 			stack = append(stack, head.Left)
	// 		}
	// 	}
	// }

}

/*后序：左右头
1、从栈中弹出一个节点cur
2、cur放入回收栈
3、先左再右放入栈
4、重复
*/
// 原因：栈弹出为头右左，回收栈弹出顺序就是栈逆序，左右头
func posOrderUnRecur(head *Node) {
	if head == nil {
		return
	}
	stack := []*Node{}
	gc_stack := []*Node{}
	// 先放入第一个结点
	stack = append(stack, head)
	for len(stack) > 0 {
		//弹出
		head = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//放入回收栈
		gc_stack = append(gc_stack, head)
		//先左后右
		if head.Left != nil {
			stack = append(stack, head.Left)
		}
		if head.Right != nil {
			stack = append(stack, head.Right)
		}
	}
	fmt.Print("后续遍历非递归：")
	for i := len(gc_stack) - 1; i > 0; i-- {
		fmt.Printf("%v-", stack[i].Val)
	}
	return
}

/*中序：左头右
1、每棵子树左树全部进栈cur，没有则下一步
2、弹出cur,打印cur
3、把cur右树进栈,没有下一步
4、重复
*/
func inOrderUnRecur(head *Node) {
	if head == nil {
		return
	}
	stack := []*Node{}
	cur := head
	// cur != nil
	for len(stack) > 0 || cur != nil {
		// 左边压栈
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			//让右子树重复左头右操作
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%v ", cur.Val)
			cur = cur.Right
		}
	}
}

func main() {
	var l1 *Node
	preOrderRecur(l1)
}
