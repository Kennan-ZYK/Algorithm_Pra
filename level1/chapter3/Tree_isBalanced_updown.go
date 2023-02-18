package chapter3

// 方法2: 自顶向下
/*
时间复杂度：O(n^2)，其中 n 是二叉树中的节点个数。
最坏情况下，二叉树是满二叉树，需要遍历二叉树中的所有节点，时间复杂度是 O(n)。
对于节点 p，如果它的高度是 d，则 height(p) 最多会被调用 d 次（即遍历到它的每一个祖先节点时）。
对于平均的情况，一棵树的高度 h 满足 O(h)=O(logn)，因为 d≤h，所以总时间复杂度为 O(nlogn)。
对于最坏的情况，二叉树形成链式结构，高度为 O(n)，
此时总时间复杂度为 O(n^2).

空间复杂度：O(n)，其中n是二叉树中的节点个数。空间复杂度主要取决于递归调用的层数，递归调用的层数不会超过n。

*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *Node) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
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
