package main

type SLinkNode struct {
	Val  int
	Next *SLinkNode
}

func S_reverse(head *SLinkNode) *SLinkNode {
	var pre *SLinkNode = nil
	var cur *SLinkNode = head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	head = pre
	return head

}

type DLinkNode struct {
	Val  interface{}
	Next *DLinkNode
	Pre  *DLinkNode
}

func Dreverse(head *DLinkNode) *DLinkNode {
	var pre *DLinkNode = nil
	var cur *DLinkNode = head

	for cur != nil {
		tmp := cur.Next
		cur.Pre = tmp
		pre = cur
		cur = tmp
	}
	head = pre
	return head
}
