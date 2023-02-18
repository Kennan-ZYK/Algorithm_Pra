/*
	链表：用快慢指针方式 ，如果快慢指针相遇，证明有环
	若相遇，让慢指针继续以1个结点速度走，快指针回到head，和慢指针一样的速度走
	当再次相遇，就是入环点，证明在笔记中

*/
package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

//无头结点
func isRing(head *LinkNode) (bool, *LinkNode) {
	if head == nil {
		return false, nil
	}
	slow, fast := head, head
	ans := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			ans = true
		}

	}
	if ans {
		fast = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		into_ring := slow
		return true, into_ring
	}
	return false, nil
}
