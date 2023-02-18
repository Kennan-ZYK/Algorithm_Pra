package chapter1

import (
	"fmt"
	"math"
)

func maxAreaOfIsland(grid [][]int) int {
	us := Init(grid)
	// fmt.Println(us)
	line, col := len(grid), len(grid[0])
	max := math.MinInt64

	for i := 0; i < line; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i-1 >= 0 && grid[i-1][j] == 1 {
				us.Union(i*col+j, (i-1)*col+j)
			}

			if j-1 >= 0 && grid[i][j-1] == 1 {
				us.Union(i*col+j, i*col+j-1)
			}
			// 这段可要可不要 因为重复判断了但是 Union中有 p==q return 所以无所谓
			// if i + 1 < line && grid[i + 1][j] == 1{
			//     us.Union(i * col + j, (i + 1) * col + j)
			// }
			// if j + 1 < col && grid[i][j+1] == 1{
			//     us.Union(i * col + j, i * col + j +1)
			// }
		}
	}
	fmt.Println(us)
	for i := 0; i < line*col; i++ {
		max = Max(max, us.count[i])
	}
	return max
}

type UnionSet struct {
	parent []int
	count  []int
}

func Init(arr [][]int) *UnionSet {
	us := &UnionSet{}
	line, col := len(arr), len(arr[0])
	us.parent = make([]int, line*col)
	us.count = make([]int, line*col)
	for i := 0; i < line; i++ {
		for j := 0; j < col; j++ {
			us.parent[i*col+j] = i*col + j
			if arr[i][j] == 1 {
				us.count[i*col+j] = 1
			} else {
				us.count[i*col+j] = 0
			}
		}
	}
	return us
}

func (us *UnionSet) FindHead(x int) int {
	if us.parent[x] != x {
		us.parent[x] = us.FindHead(us.parent[x])
	}
	return us.parent[x]
}

func (us *UnionSet) Union(n1, n2 int) {
	p, q := us.FindHead(n1), us.FindHead(n2)
	if p == q {
		return
	}
	if us.count[p] > us.count[q] {
		us.parent[q] = p
		us.count[p] += us.count[q]
	} else {
		us.parent[p] = q
		us.count[q] += us.count[p]
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
