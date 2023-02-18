package main

import "fmt"

type UnionSet struct {
	//parent[i] 表示i的父结点，根结点的父结点是本身
	parent map[int]int
	//count[i]可以用来保存本连通区域的结点个数
	count map[int]int
}

// 初始化
func InitSet(nums []int) *UnionSet {
	us := UnionSet{}
	us.parent = make(map[int]int)
	us.count = make(map[int]int)

	for _, v := range nums {
		us.parent[v] = v
		us.count[v] = 1
	}

	return &us
}

// 判断是否相同
func (us *UnionSet) isSameSet(num1, num2 int) bool {

	// 未初始化直接返回false
	_, ok1 := us.parent[num1]
	_, ok2 := us.parent[num2]
	if !ok1 || !ok2 {
		return false
	}
	return us.FindHead(num1) == us.FindHead(num2)

}

// 集合
func (us *UnionSet) Union(num1, num2 int) {
	// 未初始化直接返回
	_, ok1 := us.parent[num1]
	_, ok2 := us.parent[num2]
	if !ok1 || !ok2 {
		return
	}
	p1 := us.FindHead(num1)
	p2 := us.FindHead(num2)
	// 两个元素共同一个集合
	if p1 == p2 {
		return
	}
	// 谁长谁是头
	if us.count[num1] >= us.count[num2] {
		us.parent[num2] = num1
		us.count[num1] = us.count[num1] + us.count[num2]
	} else {
		us.parent[num1] = num2
		us.count[num2] = us.count[num1] + us.count[num2]
	}

}

func (us *UnionSet) FindHead(num int) int {

	if _, ok := us.parent[num]; !ok {
		return -1 // 根据实际情况定非法值
	}
	// 路径压缩，使元素直接指向头
	if us.parent[num] != num {
		us.parent[num] = us.FindHead(us.parent[num])
	}

	return us.parent[num]
}

func main() {
	// var n int
	// var nums []int

	// fmt.Printf("请输入几个整数的数组:")
	// fmt.Scanln(&n)
	// nums = make([]int, n)
	// cn := 0
	// for i := 0; i < n; i++ {
	// fmt.Scanf("%d\n", &nums[i])
	// }

	nums := []int{1, 2, 3, 4, 5, 6}
	// 初始化
	us := InitSet(nums)
	fmt.Println(us)
	us.Union(1, 2)
	us.Union(3, 4)
	us.Union(6, 5)
	us.Union(1, 5)
	fmt.Println(us)
	a := us.isSameSet(1, 5)
	fmt.Println(a)
}
