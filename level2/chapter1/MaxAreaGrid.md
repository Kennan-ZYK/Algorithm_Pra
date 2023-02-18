给定一个由 `0` 和 `1` 组成的非空二维数组 `grid` ，用来表示海洋岛屿地图。

一个 **岛屿** 是由一些相邻的 `1` (代表土地) 构成的组合，这里的「相邻」要求两个 `1` 必须在水平或者竖直方向上相邻。你可以假设 `grid` 的四个边缘都被 `0`（代表水）包围着。

找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 `0` 。

<img src="E:\linux学习PDF\linux\img\image-20221111151706343.png" alt="image-20221111151706343" style="zoom:100%;" />



## 并查集

```go
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
            //i * col + j ，第几行第几列，一行col个，i行
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

```



分析：用并行的算法怎么处理？

平分后，左半边和右半边同时运行，然后对平分线左右两侧进行判断，左右有1 证明此处，左右应该联通，union即可

```go
0   1   2    3   4   5   6  20  8  9  10 11 12 	
0   0   1    0   0   0   0  1   0  0  0   0  0 
        
13  14  15   16  17 18   19 20 20 20  23 24 25 
0    0  0    0   0  0    0  4  1  1   0   0  0  

26  40 40   29  43  31   32 33 34 35  36 37 38  
0   1   1   0   1   0     0  0  0  0   0  0  0  

39  40 41   42  43  43   45 46 60 48 60 50 51  
0   4   0   0   5    1    0  0   1 0  1  0  0  

52  40  54 55   43  43   58  59 60 60 60 63 64 
0    1   0  0    1   1    0   0  6  1  2  0  0

65  66 67  68  69   70   71 72  73 74 60 76 77 
0    0  0  0    0   0    0  0  0   0  1  0  0 

78  79 80  81  82   83   84 98 98 98  88 89  90 
0   0  0   0    0   0    0  1   1  1  0   0  0 

91  92 93  94  95  96    97 98 98 100 101 102 103
0   0  0   0   0    0    0   5  1  0   0   0   0 

```

