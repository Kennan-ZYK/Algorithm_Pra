package chapter3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(process(root.Left)-process(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

// 求深度
func process(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(process(root.Left), process(root.Right)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
