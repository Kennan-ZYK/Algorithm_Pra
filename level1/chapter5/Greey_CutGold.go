package main

import (
	"container/heap"
	"fmt"
)

/*
一块金条切成两半，是需要花费和长度数值一样的铜板的。比如长度为20的金条，不管切成长度多大的两半，都要花费20个铜板。
一群人想整分整块金条，怎么分最省铜板?
例如,给定数组{10,20,30}，代表一共三个人，整块金条长度为10+20+30=60。金条要分成10,20,30三个部分。
如果先把长度60的金条分成10和50，花费60;再把长度50的金条分成20和30，花费50;一共花费110铜板。
但是如果先把长度60的金条分成30和30，花费60﹔再把长度30金条分成10和20,花费30;一共花费90铜板。
输入一个数组，返回分割的最小代价。
*/

/*
思路：用哈夫曼树的思想，为什么用哈夫曼树呢？
因为哈夫曼树又叫最优二叉树，构建的树的带权路径长度最小
思想就是让每个人就代表一个权，构建哈弗曼树
做法：
用小根堆来保存每个人分的的金条，也就是权值，
连续弹出2次堆顶元素，进行相加，再把相加后的值放回小根堆，
继续重复以上操作，直到堆为空
*/

//小根堆
type IntHeap []int

func (c IntHeap) Len() int           { return len(c) }
func (c IntHeap) Less(i, j int) bool { return c[i] < c[j] }
func (c IntHeap) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c *IntHeap) Push(x interface{}) {
	*c = append(*c, x.(int))

}
func (c *IntHeap) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}
func lessMoney(arr []int) int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return 0
	}
	h := IntHeap(arr)
	heap.Init(&h)
	ans := 0
	twosum := 0
	for i := 1; i < arrLen; i++ {
		twosum = heap.Pop(&h).(int) + heap.Pop(&h).(int)
		ans += twosum
		heap.Push(&h, twosum)
	}
	return ans
}
func main() {
	arr := []int{10, 30, 20}
	ret := lessMoney(arr)
	fmt.Println(ret)
}
