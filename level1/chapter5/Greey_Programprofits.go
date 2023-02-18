package chapter5

import "container/heap"

/*
输入: 正数数组costs、正数数组profits、正数K、正数M
costs[i]表示i号项目的花费
profits[i]表示i号项目在扣除花费之后还能挣到的钱(利润)
K表示你只能串行的最多做k个项目
M表示你初始的资金
说明: 每做完一个项目，马上获得的收益，可以支持你去做下一个项目。不能并行的做项目。
输出：你最后获得的最大钱数。
*/

/*
思路：用小根堆保存以花费准排序，再用大根堆保存以利润为准的排序
*/

type Pair struct {
	costs   int
	profits int
}

type IntHeap struct {
	heap []Pair
	//true是小根堆，false是大根堆
	bool
}

func (h IntHeap) Len() int { return len(h.heap) }
func (h IntHeap) Less(i, j int) bool {
	if h.bool {
		return h.heap[i].costs < h.heap[j].costs
	} else {
		return h.heap[i].profits > h.heap[j].profits
	}

}

// 小根堆  > 大根堆
func (h IntHeap) Swap(i, j int) { h.heap[i], h.heap[j] = h.heap[j], h.heap[i] }

func (h *IntHeap) Push(x interface{}) {
	h.heap = append(h.heap, x.(Pair))
}

func (h *IntHeap) Pop() interface{} {
	old := h
	n := len(old.heap)
	x := old.heap[n-1]
	h.heap = old.heap[0 : n-1]
	return x

}

// K表示你只能串行的最多做k个项目
// M表示你初始的资金
func findMaximizedCapital(K, M int, Profixs, Costs []int) int {
	c := &IntHeap{bool: true}
	heap.Init(c)
	p := &IntHeap{bool: false}
	heap.Init(p)
	for i := 0; i < len(Costs); i++ {
		cur := Pair{Costs[i], Profixs[i]}
		heap.Push(c, cur)
	}
	// fmt.Println(c.heap)
	for i := 0; i < K; i++ {
		for len(c.heap) > 0 && c.heap[0].costs <= M {
			top := heap.Pop(c)
			heap.Push(p, top)
		}
		if len(p.heap) == 0 {
			return M
		} else {
			down := heap.Pop(p).(Pair)
			M += down.profits
		}
	}
	return M
}

func main() {

}
