package chapter2

import "container/heap"

/*
由一个代表题目，引出一种结构[题目]
有一个整型数组arr和一个大小为w的窗口从数组的最左边滑到最右边，窗口每次 向右边滑一个位置。
例如，数组为[4,3,5,4,3,3,6,7]，窗口大小为3时:
[4 3 5]4 3 3 6 7
4[3 5 4]3 3 6 7
4 3[5 4 3]3 6 7
4 3 5[4 3 3]6 7
4 3 5 4[3 3 6]7
4 3 5 4 3[3 6 7]
窗口中最大值为5 窗口中最大值为5 窗口中最大值为5 窗中最大值为4 窗口中最大值为6窗口中最大值为7
如果数组长度为n，窗口大小为w，则一共产生n-w+1个窗口的最大值。请实现一个函数。
输入:整型数组arr，窗口头小为w。
输出:一个长度为n-w+1的数组res，res[i]表示每一种窗口状态下的 以本题为例，结果应该
返回{5,5,5,4,6.7}。
*/

/*
思路: 准备一个双端队列，头部尾部都可以进出结点，存放数组下标，下标即可知道位置，也可以直到值
双端队列必须是由大到小排序，发现当前数大于尾部值，弹出尾部值，直到发现前面值大于自己或者清空了填入，弹出数据永远不找回
*/

func maxSlidingWindow(nums []int, k int) []int {
	Queue := []int{}
	ans := []int{}

	for right := 0; right < len(nums); right++ {
		for len(Queue) != 0 && nums[right] >= nums[Queue[len(Queue)-1]] {
			Queue = Queue[:len(Queue)-1]
		}
		// 放入下标
		Queue = append(Queue, right)

		left := right - k + 1
		// 过期
		for Queue[0] < left {
			Queue = Queue[1:]
		}
		// 把队列头(最大值)放入 ans，
		if right >= k-1 {
			ans = append(ans, nums[Queue[0]])
		}
	}

	return ans
}

// 方法2

func maxSlidingWindow2(nums []int, k int) []int {

	var res []int
	var t = wi{}
	for i := 0; i < len(nums); i++ {
		heap.Push(&t, w{i, nums[i]})
		if len(t) >= k {
			for t[0].index <= i-k {
				heap.Pop(&t)
			}
			res = append(res, t[0].num)
		}
	}
	return res
}

type w struct{ index, num int }
type wi []w

func (h wi) Len() int            { return len(h) }
func (h wi) Less(i, j int) bool  { return h[i].num > h[j].num }
func (h wi) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *wi) Push(v interface{}) { *h = append(*h, v.(w)) }
func (h *wi) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
