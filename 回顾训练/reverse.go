package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func isreverse(head *LinkNode) bool {
	if head == nil {
		return true
	}
	f, s := head, head
	for f.Next != nil || f.Next.Next != nil {
		f = f.Next.Next
		s = s.Next.Next
	}
	// 此时s是中点，若是回文左右两边要相等
	f = s.Next
	s.Next = nil
	for f != nil {
		f.Next = s
		s = f
		f = f.Next
	}
	f = head
	// 此时s为右版区的头
	for f != s {

	}
}
