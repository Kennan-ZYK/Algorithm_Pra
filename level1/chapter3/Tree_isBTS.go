package chapter3

import "math"

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

/*二叉搜索树特点：
若它的左子树不为空，则左子树上所有节点的值都小于根节点的值
若它的右子树不为空，则右子树上所有节点的值都大于根节点的值
它的左右子树也分别为二叉搜索树
*/

/******判断是否为二叉搜索树（二叉排序树）中序遍历即可******/

// 方法1

// 父亲结点
var preVal = -1

func isBTS(head *Node) bool {
	// 空树或子结点
	if head == nil {
		return true
	}
	var isLeftBts bool = isBTS(head.Left)
	if !isLeftBts {
		return false
	}
	if head.Val <= preVal {
		return false
	} else {
		preVal = head.Val
	}
	// 右子树判断
	return isBTS(head.Right)
}

// 方法2
func isBST(root *Node) bool {
	if root == nil {
		return true
	}
	res := []int{}
	process(root, &res)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

func process(root *Node, res *[]int) *[]int {
	if root == nil {
		return &[]int{}
	}
	process(root.Left, res)
	(*res) = append(*res, root.Val)
	process(root.Right, res)
	return res
}

//方法3
//非递归
func isBTS3(head *Node) bool {
	if head == nil {
		return true
	}
	// 保存上一结点
	preVal := math.MinInt64
	Stack := []*Node{}
	for head != nil || len(Stack) > 0 {
		if head.Left != nil {
			Stack = append(Stack, head)
			head = head.Left
		} else {
			head = Stack[len(Stack)-1]
			Stack = Stack[:len(Stack)-1]
			if head.Val <= preVal {
				return false
			} else {
				preVal = head.Val
			}
			// 遍历右子树
			head = head.Right
		}
	}
	return true
}

// leetcode题解
func isValidBST(root *Node) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *Node, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
