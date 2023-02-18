package chapter3

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

/*
平衡二叉树定义：
每个节点的左子树和右子树的高度差至多等于 1
*/

// 方法1: 自底向上(最优)
/*
时间复杂度：O(n)
空间复杂度：O(n)
*/
func isBalanced(root *Node) bool {
	return height(root) >= 0
}

func height(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
