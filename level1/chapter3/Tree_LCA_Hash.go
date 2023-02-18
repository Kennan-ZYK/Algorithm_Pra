package chapter3

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func lowestCommonAncestor(head, o1, o2 *Node) *Node {
	fatherMap := map[*Node]*Node{}
	fatherMap[head] = head
	process(head, fatherMap)
	set := map[*Node]bool{}
	// cur来到o1位置
	cur := o1
	// 只有根结点的父结点是自己
	// 保存cur之上所有结点
	for cur != fatherMap[cur] {
		set[cur] = true
		cur = fatherMap[cur]
	}
	set[head] = true

	cur = o2
	for cur != fatherMap[cur] {
		// 先判断是否在里面，不在再往上窜
		// 不先判断会造成少比较本身一次  如 [-1,0,3,-2,4,null,null,8] o1=8 o2=0
		if _, ok := set[cur]; ok {
			return cur
		}
		cur = fatherMap[cur]
	}
	// 防止o2为根结点，就不判断  如：[1,2] o1=2,o2=1
	if _, ok := set[cur]; ok {
		return cur
	}
	return nil

}

// 求父亲
func process(head *Node, fatherMap map[*Node]*Node) {
	if head == nil {
		return
	}
	fatherMap[head.Left] = head
	fatherMap[head.Right] = head
	process(head.Left, fatherMap)
	process(head.Right, fatherMap)
	return
}
