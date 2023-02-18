/*
【题目】给定两个可能有环也可能无环的单链表，头节点head1和head2。请实现一个函数，如果两个链表相交，请返回相交的第一个节点。如果不相交，返回null
【要求】如果两个链表长度之和为N,时间复杂度请达到0(N)，额外空间复杂度请达到0(1)。
*/

/*
情况1：loop1 , loop2 ==nil
情况2：loop1 , loop2 一个为空，一个不为空  （一定不相交）
情况3：loop1,loop2都有环
	   1) 各自独立，判断end
	   2) 入环结点一致 getNode()得出 loop1 = loop2 ,然后以入环结点为终点，做和无环相交一样的方法
	   3) 各自都有入环结点 让h1链表继续走，如果回到loop1前，遇到loop2，则相交，相交结点loop1,loop2都是
*/
// 时间复杂度O(N)
package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

//无头结点
//获取入环结点
func getloopNode(head *LinkNode) *LinkNode {
	//为了判断快慢指针是否能用
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for slow != fast {
		//长度奇数偶数情况
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

//无环链表
func noLoop(h1 *LinkNode, h2 *LinkNode) *LinkNode {
	if h1 == nil || h2 == nil {
		return nil
	}
	cur1, cur2 := h1, h2
	//长度
	n := 0
	for cur1 != nil {
		n++
		cur1 = cur1.Next
	}
	for cur2 != nil {
		n--
		cur2 = cur2.Next
	}
	//end不相等直接结束
	if cur1 != cur2 {
		return nil
	}
	//长为h1，短为h2
	if n > 0 {
		cur1 = h1
		cur2 = h2
	} else {
		cur1 = h2
		cur2 = h1
	}
	n = abs(n)
	//让cur1走完差值
	for n != 0 {
		n--
		cur1 = cur1.Next
	}
	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return cur1
}

//有环链表
func bothLoop(h1, loop1, h2, loop2 *LinkNode) *LinkNode {
	if h1 == nil || h2 == nil {
		return nil
	}
	cur1, cur2 := h1, h2
	//计数器
	n := 0
	//情况2
	if loop1 == loop2 {
		for cur1 != loop1 {
			n++
			cur1 = cur1.Next
		}
		for cur2 != loop2 {
			n--
			cur2 = cur2.Next
		}
		// 比较长度
		if n > 0 {
			cur1, cur2 = h1, h2
		} else {
			cur1, cur2 = h2, h1
		}
		n = abs(n)
		// 走差值
		for n != 0 {
			n--
			cur1 = cur1.Next
		}
		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else {
		cur1 = loop1.Next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.Next
		}
		return nil
	}

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getIntersectionNode(h1, h2 *LinkNode) *LinkNode {
	if h1 == nil || h2 == nil {
		return nil
	}
	var loop1, loop2 *LinkNode
	loop1 = getloopNode(h1)
	loop2 = getloopNode(h2)
	// 情况1
	if loop1 == nil && loop2 == nil {
		return noLoop(h1, h2)
	}
	// 情况3
	if loop1 != nil && loop2 != nil {
		return bothLoop(h1, loop1, h2, loop2)
	}
	// 情况2
	return nil
}
