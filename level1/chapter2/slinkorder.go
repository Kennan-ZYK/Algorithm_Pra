package singleLink

//【题目】给定一个单链表的头节点head，节点的值类型是整型，再给定一个整数pivot。
// 实现一个调整链表的函数，将链表调整为左部分都是值小于pivot的节点，中间部分都是值等于pivot的节点，右部分都是值大于pivot的节点。

type LinkNode struct {
	Val  int
	Next *LinkNode
}

//链表含头指针
func getNode(head *LinkNode) []*LinkNode {
	if head.Next == nil {
		return []*LinkNode{}
	}
	arr := []*LinkNode{}
	i := 0
	cur := head.Next
	for cur != nil {
		arr[i] = cur
		cur = cur.Next
	}
	return arr
}

func partition(arr []LinkNode, point int) []int {
	if arr == nil {
		return []int{}
	}
	//设置左右边界
	left, right := -1, len(arr)
	i := 0
	for i < right {
		if arr[i].Val < point {
			left++
			swap(arr, i, left)
			i++
		} else if arr[i].Val > point {
			right--
			swap(arr, i, right)
		} else {
			i++
		}
	}
	return []int{left, right}
}

func recoverLink(arr []LinkNode, head *LinkNode) {
	if arr == nil {
		return
	}
	length := len(arr)
	cur := head.Next
	// newnode := &LinkNode{Val: -1, Next: nil}
	for i := 0; i < length; i++ {
		cur.Val = arr[i].Val
	}

}

func swap(arr []LinkNode, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
