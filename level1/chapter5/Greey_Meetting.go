package chapter5

import (
	"fmt"
	"sort"
)

/*
题目：一些项目要占用一个会议室宣讲，会议室不能同时容纳两个项目的宣讲。
给你每一个项目开始的时间和结束的时间(给你一个数组，里面是一个个具体的项目)，
你来安排宣讲的日程，要求会议室进行的宣讲的场次最多。
返回这个最多的宣讲场次。
*/
/*
为什么按结束时间呢？
局部最优，当 (1-4，3-5，0-6) 这3个时间点，选择1-4，因为在这3个时间段内，选谁都是只能选一个，
在这个局部内，选最早结束，后面还有会议的话，反而能安排更多会议


思路：先按结束时间从小到大排序，如果结束时间相同，就按开始时间从大到小排序（保证持续时间最短），
对排好序的进行遍历，timePoint为一天项目最早开始时间，
从大于等于timePoint的会议开始安排，选择完一个会议就更新timePoint为会议的结束时间，直到遍历完
*/

type Program struct {
	start int
	end   int
}
type ProgramSlice []Program

func (p ProgramSlice) Len() int { // 重写 Len() 方法
	return len(p)
}

func (p ProgramSlice) Swap(i, j int) { // 重写 Swap() 方法
	p[i], p[j] = p[j], p[i]
}
func (p ProgramSlice) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	// 结束时间相同，按开始时间大到小的排序，（确保持续时间最短）
	if p[i].end == p[j].end {
		return p[i].start > p[j].start
	}
	return p[i].end < p[j].end
}

func MaxMeeting(program []Program, timePoint int) int {
	sort.Sort(ProgramSlice(program))
	res := 0
	for i := 0; i < len(program); i++ {
		if timePoint <= program[i].start {
			res += 1
			timePoint = program[i].end
		}
	}
	return res
}

/**方法2**/
func main() {
	programs := make([]*Program, 3)
	programs[0] = &Program{1, 27}
	programs[1] = &Program{26, 31}
	programs[2] = &Program{29, 300}
	ret := bestArrange(programs)
	fmt.Println(ret)
}

// 会议的开始时间和结束时间，都是数值，不会 < 0
func bestArrange(programs []*Program) int {
	//按会议的结束时间排序
	sort.SliceIsSorted(programs, func(i, j int) bool {
		return programs[i].end < programs[j].end
	})
	programsLen := len(programs)
	timeLine := 0
	result := 0
	// 依次遍历每一个会议，结束时间早的会议先遍历
	for i := 0; i < programsLen; i++ {
		if timeLine <= programs[i].start {
			result++
			timeLine = programs[i].end
		}
	}
	return result
}
