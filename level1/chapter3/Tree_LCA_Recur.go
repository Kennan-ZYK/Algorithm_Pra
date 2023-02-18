package chao1ter3

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

// 两结点最小公共祖先
/*
两种情况：
1、o1是o2的LCA 或者 o2是o1的LCA
2、o1,o2互不为LCA,要向上找LCA
画图好理解！！！
*/
func lowestCommonAncestor(head, o1, o2 *Node) *Node {
	if head == nil || head == o1 || head == o2 {
		return head
	}
	left := lowestCommonAncestor(head.Left, o1, o2)
	right := lowestCommonAncestor(head.Right, o1, o2)
	// 情况2 左右都不为空则找到了
	if left != nil && right != nil {
		return head
	}
	if left != nil {
		return left
	}
	return right
}
