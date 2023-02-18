package chapter5

import "container/heap"

/*
题目:一个数据流中，随时获取中位数
*/

/*
中位数：在排好序的数列中，中间值，偶数情况是((len/2) + (len/2)+1) /2,奇数len/2
思路:利用两个堆，一个大根堆，一个小根堆。要满足以下条件
1、把要准备放入堆中的数cur，与大根堆的顶比较，
cur <= 大根堆顶，放入大根堆，否则放入小根堆，
2、当大根堆和小根堆中数的数量差值相差 2时，把数量大的堆顶弹出进入另外一个堆
最后会形成 较小的一半数在大根堆，较大的一半数在小根堆
*/

/*
 1. 数组
未排序的数组：采用快排的 Partition()。因此，插入时间复杂度是O(1) ，而找中位数的时间复杂度是 O(n)。
排序的数组：适合使用插入排序，插入的时间复杂度是 O(n)，而找中位数的时间复杂度是 O(1)。
2.链表
排序的链表，插入的时间复杂度还是O(n) ，找中位数的时间复杂度是 O(n)，这里对链表中间节点做个标记，可以降低到O(1)
3.二叉搜索树BST
插入的时间复杂度平均O(log(n))，找到中位数的时间复杂度是 O(log(n)),最坏情况退化为链表O(n)。
4.平衡二叉树树AVL
二叉搜索树的改进，插入数据的时间复杂度不变，找到中位数的时间复杂度是O(1)。
5.堆
插入数据的时间复杂度是 O(log(n))。而中位数肯定在两个堆的堆顶元素中，找到中位数的时间复杂度是 O(1)。
*/
type IntHeap struct {
	heap []int
	// true 为小根堆，false 为大根堆
	bool
}

func (h IntHeap) Len() int      { return len(h.heap) }
func (h IntHeap) Swap(i, j int) { h.heap[i], h.heap[j] = h.heap[j], h.heap[i] }
func (h IntHeap) Less(i, j int) bool {
	if h.bool {
		return h.heap[i] < h.heap[j]
	}
	return h.heap[i] > h.heap[j]
}
func (h *IntHeap) Push(x interface{}) {
	(*h).heap = append((*h).heap, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old.heap)
	x := old.heap[n-1]
	h.heap = h.heap[:n-1]
	return x
}

func (h *IntHeap) Top() int {
	return h.heap[0] // 注意：在这里(golang里)堆顶（无论是大顶堆，还是小顶堆）是第1个元素
}

func GetMid(arr []int) int {
	// 初始化大小根堆
	hBig := &IntHeap{bool: false}
	heap.Init(hBig)
	hSmall := &IntHeap{bool: true}
	heap.Init(hSmall)
	// 先把第一个放入大根堆
	heap.Push(hBig, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] <= hBig.heap[0] {
			hBig.Push(arr[i])
		} else {
			hSmall.Push(arr[i])
		}
		if hBig.Len()-hSmall.Len() == 2 {
			top := hBig.Pop().(int)
			hSmall.Push(top)
		} else if hSmall.Len()-hBig.Len() == 2 {
			top := hSmall.Pop().(int)
			hBig.Push(top)
		}
	}

	topb := hBig.Pop().(int)
	tops := hSmall.Pop().(int)
	if len(arr)%2 == 0 {
		return (topb + tops) / 2
	} else {
		if hBig.Len() > hSmall.Len() {
			return tops
		} else {
			return topb
		}
	}
}
