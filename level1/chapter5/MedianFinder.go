package chapter5

import "container/heap"

// 数据流中位数

type MedianFinder struct {
	hBig   *IntHeap // 前半部分，用大顶堆存储，都是较小的数（都小于后半部分）
	hSmall *IntHeap
}

func Constructor() MedianFinder {
	hBig := &IntHeap{bool: false}
	heap.Init(hBig)
	hSmall := &IntHeap{bool: true}
	heap.Init(hSmall)
	return MedianFinder{hBig, hSmall}
}

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

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.hBig, num) // 先加入到大顶堆

	heap.Push(this.hSmall, heap.Pop(this.hBig).(int)) // 调整两个堆平衡，此时从大顶堆Pop出最大元素，加入到小顶堆。
	for this.hBig.Len() < this.hSmall.Len() {
		heap.Push(this.hBig, heap.Pop(this.hSmall).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.hBig.Len()+this.hSmall.Len())%2 == 0 {
		return float64(this.hBig.Top()+this.hSmall.Top()) * 0.5
	} else {
		// fmt.Println(this.hBig.Len() , this.hSmall.Len())
		if this.hBig.Len() > this.hSmall.Len() {
			return float64(this.hBig.Top())
		} else {
			return float64(this.hSmall.Top())
		}
	}

}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
