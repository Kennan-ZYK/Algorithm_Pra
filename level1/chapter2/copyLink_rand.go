/*
题目：
rand指针是单链表节点结构中新增的指针，rand可能指向链表中的任意一个节点，也可能指向null。
给定一个由Node节点类型组成的无环单链表的头节点head，请实现一个函数完成这个链表的复制，并返回复制的新链表的头节点。
【要求】时间复杂度O(N)，额外空间复杂度0(1)
*/
package copyLink_rand

type LinkNode struct {
	Val  int
	Next *LinkNode
	Rand *LinkNode
}

//含头结点
func copyrand(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}
	cur := head.Next
	//copy老元素
	for cur != nil {
		old := cur.Next //老元素
		newcode := new(LinkNode)
		newcode.Val = cur.Val

		cur.Next = newcode //新元素
		newcode.Next = old //1->1'->2
		cur = old          // 完成一对,下一队
	}
	cur = head.Next
	var copynode *LinkNode = nil
	//让copy元素去找老元素的rand
	for cur != nil {
		tmp := cur.Next.Next //老元素
		copynode = cur.Next  //新元素
		if cur.Rand != nil {
			copynode.Rand = cur.Rand.Next
		} else {
			copynode.Rand = nil
		}
		cur = tmp
	}
	//拆分，让copy和老元素分开
	cur = head.Next
	res := cur.Next //保存copy链表开头，以便返回
	for cur != nil {
		old := cur.Next.Next
		copynode = cur.Next
		cur.Next = old
		if copynode.Next != nil {
			copynode.Next = old.Next
		} else {
			copynode.Next = nil
		}
		cur = old
	}

	return res
}
