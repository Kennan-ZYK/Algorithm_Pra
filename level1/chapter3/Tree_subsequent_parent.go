package chapter3

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// 后继结点(中序遍历后的结点)
// 如：213  1为2的后继结点

/*
两种情况:
1. x 有右子树，找下一结点的左子树，若没有左子树,则返回右子树根结点
2. x 没有右子树，往上找父节点，只要当前结点是父结点的左结点，则返回
*/

func getSubsequent(node *Node) *Node {
	if node == nil {
		return node
	}
	// 情况1： 右树不为空
	if node.Right != nil {
		return getLetMost(node.Right)
	} else {
		// 情况2：
		parent := node.Parent
		// 当前结点为父结点的有孩子
		for parent != nil && parent.Left != node {
			node = parent
			parent = node.Parent
		}
		return parent
	}
}

func getLetMost(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	// 若没有左子树，根据中序遍历，左头右，那么当前结点是头结点，返回即可
	return node
}
