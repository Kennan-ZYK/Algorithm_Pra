// 在Golang的container/heap包下面有堆的接口，只需要实现接口即可简单的实现二叉堆。
// 如下为该接口的代码，显然，实现heap接口的同时，得先实现sort接口。
package chapter1

import (
	"container/heap"
	"fmt"
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

// go标准文档：
// 任何实现了本接口的类型都可以用于构建最小堆。最小堆可以通过heap.Init建立，数据是递增顺序或者空的话也是最小堆。最小堆的约束条件是：
// !h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
// 注意：接口的Push和Pop方法是供heap包调用的，请使用heap.Push和heap.Pop来向一个堆添加或者删除元素。

type MyHeap []int

func (m MyHeap) Len() int {
	return len(m)
}

//如果实现大顶堆，则需要调整一下Less函数
func (m MyHeap) Less(i, j int) bool {
	if m[i] < m[j] {
		return true
	}
	return false
}

func (m MyHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MyHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MyHeap) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func main() {
	myHeap := MyHeap{3, 6, 2, 1, 4, 5}
	heap.Init(&myHeap) //如果没有初始值，可以不需要使用此方法
	heap.Push(&myHeap, 7)
	fmt.Println(myHeap) //[1 3 2 6 4 5 7]
	pop := heap.Pop(&myHeap)
	fmt.Println(pop)    //1
	fmt.Println(myHeap) //[2 3 5 6 4 7]
}
