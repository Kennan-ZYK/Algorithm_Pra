/*
题目：给定一个单链表的头结点head，判断该链表是否为回文结构
要求：链表长度为N，时间复杂度为O(N),额外空间复杂度为O(1).
*/

/*方法:采用快慢指针，当 fast 到结尾时，slow 刚好指向终点(奇数)/终点前一位(偶数)

 */
package main

type Node struct {
	Val  int
	Next *Node
}

//head 是头结点
func isPalindrome(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast := head.Next
	slow := head.Next
	//寻找mid
	for fast.Next != nil && fast.Next.Next != nil { //tips:这里只管找mid，fast不一定是终点位置
		fast = fast.Next.Next
		slow = slow.Next
	}
	//让mid右边反转
	pre := slow.Next
	slow.Next = nil
	var tmp *Node = nil
	for pre != nil {
		tmp = pre.Next
		pre.Next = slow
		slow = pre
		pre = tmp
	}
	//slow此时为最后一位，用tmp记录，以便复原
	tmp = slow

	//让fast为右head，slow为左head
	fast = slow
	slow = head
	ans := true
	for fast != slow {
		if slow.Val != fast.Val {
			ans = false
			break
		}
		slow = slow.Next
		fast = fast.Next
	}

	//恢复链表
	fast = tmp.Next
	tmp.Next = nil
	for fast != nil {
		slow = fast.Next
		fast.Next = tmp
		tmp = fast
		fast = slow
	}

	return ans

}
