package expend

// 大小根堆实现
// 但是如果我们同时需要大根堆和小根堆，难道要实现两遍或定义两次吗？实际上由于heap提供了足够灵活的接口，我们可以根据自己的需求做一些更改，也比如自己封装结构体实现优先队列等等。这里我们展示如何用一个数据结构根据属性值完成大根堆和小根堆的功能

type IntHeap struct {
	heap []int
	//true是小根堆，false是大根堆
	bool
}

func (h IntHeap) Len() int { return len(h.heap) }
func (h IntHeap) Less(i, j int) bool {
	if h.bool {
		return h.heap[i] < h.heap[j]
	} else {
		return h.heap[i] > h.heap[j]
	}

}                               // 小根堆  > 大根堆
func (h IntHeap) Swap(i, j int) { h.heap[i], h.heap[j] = h.heap[j], h.heap[i] }

func (h *IntHeap) Push(x interface{}) {
	h.heap = append(h.heap, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := h
	n := len(old.heap)
	x := old.heap[n-1]
	h.heap = old.heap[0 : n-1]
	return x

}

// 当bool是true时，是一个小根堆，否则是大根堆。
