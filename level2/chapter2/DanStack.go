package main

import "fmt"

/*
在数组中想找到一个数，左边和右边比这个数大、且离这个数最近的位置
如果对每一个数都想求这样的信息，能不能整体代价达到0(N)?
需要使用到单调栈结构单调栈结构的原理和实现
5 4 6 7 2 3 0 1
	index->value
    左      右
5   无     2->6
4   0->5   2->6
6   无     3->7
7   无     无
2   3->7   5->3
3   3->7   无
0   5->3   7->1
1   5->3   无
*/

// 无重复值
// 下标->[左下标，右下标]
func Findnums1(nums []int) map[int][]int {
	stack := []int{}
	ans := make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			preIndex := stack[len(stack)-1]
			ans[preIndex] = make([]int, 2)
			// 左,如果只剩一个是没有左边值的
			if len(stack) == 1 {
				stack = stack[:len(stack)-1]
				ans[preIndex][0] = -1
			} else {
				// 先弹出自己，在记录自己下一个元素为左边界
				stack = stack[:len(stack)-1]
				ans[preIndex][0] = stack[len(stack)-1]
			}
			// 右
			ans[preIndex][1] = i
		}

		stack = append(stack, i)
	}
	for len(stack) != 0 {
		preIndex := stack[len(stack)-1]
		// 记录长度
		cnt := len(stack)
		ans[preIndex] = make([]int, 2)
		// 右边无大于值
		ans[preIndex][1] = -1
		stack = stack[:len(stack)-1]
		// 左边
		if cnt == 1 {
			ans[preIndex][0] = -1
		} else {
			ans[preIndex][0] = stack[len(stack)-1]
		}

	}
	return ans
}

// 有重复值 map[int][]int

type Link struct {
	index int
	next  *Link
}

func TopVal(head *Link) int {
	return head.index
}

func Findnums2(nums []int) map[int][]int {
	stack := []*Link{}
	// 下标->[左，右]
	ans := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		for len(stack) != 0 && nums[i] > nums[TopVal(stack[len(stack)-1])] {
			preIndex := TopVal(stack[len(stack)-1])
			ans[preIndex] = make([]int, 2)
			// 当只有一个说明没有左大值，只有有大值
			if len(stack) == 1 {
				// 逐一对链表进行弹出
				for stack[len(stack)-1].next != nil {
					stack[len(stack)-1] = stack[len(stack)-1].next
					ans[preIndex][0] = -1
					ans[preIndex][1] = i
					// 更新，因为之前记录的是链表第一个元素
					preIndex = TopVal(stack[len(stack)-1])
				}
				// 没有左大值
				ans[preIndex][0] = -1
				ans[preIndex][1] = i
				stack = stack[:len(stack)-1]
			} else {
				// 逐一对链表进行弹出
				for stack[len(stack)-1].next != nil {
					stack[len(stack)-1] = stack[len(stack)-1].next
					ans[preIndex][0] = TopVal(stack[len(stack)-2])
					ans[preIndex][1] = i
					// 更新，因为之前记录的是链表第一个元素
					preIndex = TopVal(stack[len(stack)-1])
				}
				// 没有后续结点
				ans[preIndex][0] = TopVal(stack[len(stack)-2])
				ans[preIndex][1] = i
				stack = stack[:len(stack)-1]
			}

		}
		if len(stack) != 0 && nums[i] == nums[TopVal(stack[len(stack)-1])] {
			link := &Link{i, nil}
			stack[len(stack)-1].next = link
		}
		link := &Link{i, nil}
		stack = append(stack, link)
	}

	for len(stack) != 0 {
		preIndex := TopVal(stack[len(stack)-1])
		if len(stack) == 1 {
			for stack[len(stack)-1].next != nil {
				ans[preIndex] = make([]int, 2)
				ans[preIndex][0] = -1
				ans[preIndex][1] = -1
				stack[len(stack)-1] = stack[len(stack)-1].next
				// 更新，因为之前记录的是链表第一个元素
				preIndex = TopVal(stack[len(stack)-1])
			}
			ans[preIndex] = make([]int, 2)
			ans[preIndex][0] = -1
			ans[preIndex][1] = -1

		} else {
			for stack[len(stack)-1].next != nil {
				ans[preIndex] = make([]int, 2)
				ans[preIndex][0] = TopVal(stack[len(stack)-2])
				ans[preIndex][1] = -1
				stack[len(stack)-1] = stack[len(stack)-1].next
				// 更新，因为之前记录的是链表第一个元素
				preIndex = TopVal(stack[len(stack)-1])
			}
			ans[preIndex] = make([]int, 2)
			ans[preIndex][0] = TopVal(stack[len(stack)-2])
			ans[preIndex][1] = -1
		}
		stack = stack[:len(stack)-1]
	}
	return ans
}

func main() {
	nums := []int{5, 4, 6, 7, 2, 2, 3, 0, 1}
	// Findnums2(nums)
	// a := Findnums1(nums)
	b := Findnums2(nums)
	fmt.Println(b)
}
