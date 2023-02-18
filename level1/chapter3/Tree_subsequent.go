package chapter3

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// leetcode 后继者
//
func inorderSuccessor(root, p *Node) *Node {
	st := []*Node{}
	var pre, cur *Node = nil, root
	for len(st) > 0 || cur != nil {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		if pre == p {
			return cur
		}
		pre = cur
		cur = cur.Right
	}
	return nil
}
