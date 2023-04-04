package main

import (
	_ "errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Link struct {
	index int
	next  *Link
}
type Edge struct {
	weight int
	from   *Node
	to     *Node
}
type unionSet struct {
	//parent[i] 表示i的父结点，根结点的父结点是本身
	parent map[interface{}]int
	//count[i]可以用来保存本连通区域的结点个数
	count map[int]int
}
type Node struct {
	pass *Node
	end  map[byte]*Node
}
type StringSlice []string

func (x StringSlice) Len() int           { return len(x) }
func (x StringSlice) Less(i, j int) bool { return x[i]+x[j] < x[j]+x[i] }
func (x StringSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func a(arr *[]int) (c []int) {
	// fmt.Printf("%v", a)
	return c
}

func generateRandomArry(lenght int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, int(rand.Float32()*float32(lenght))+1)
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		// +1 是因为rand.Float32[0,x),得加1才能得到x
		arr[i] = int(rand.Float32()*float32(max)) + 1
	}
	return arr
}
func DeferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

type IntSlice []int

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// func Int63n(n int64) int64 {}
func main() {
	// a := []int{2, 1, 451, 2}
	// tmp := make(map[int]int)
	// for _, v := range a {
	// 	tmp[v]++
	s := []int{3, 2, 1, 4}
	// s = append(s, 2)
	// }
	// for i := 0; i < len(tmp)+1; i++ {
	// 	fmt.Print(tmp[i])
	// }

	// h = IntSlice(a)
	// sort.Sort(IntSlice(a))
	fmt.Println(s[:3])
	// a := new(Graph)
	// ts := fmt.Sprintf("%x", time.Now().Unix())
	// t1, _ := strconv.ParseInt(ts, 16, 32)
	// // log.Println(ts)
	// time.Sleep(1 * time.Second)
	// ty := fmt.Sprintf("%x", time.Now().Unix())
	// t2, _ := strconv.ParseInt(ty, 16, 32)
	// // log.Println(ty)
	// ty := fmt.Sprintf("%x", time.Now().Unix())
	// fmt.Println(ty[:8])
	// fmt.Println((math.Ceil(float64(138604) / (5 * 1024 * 1024))))
	// fmt.Println(float64(138604) / 5 * 1024)
	// math.Ceil((5 * 1024 * 1024)
	// log.Println(time.Now().Unix())
	// log.Println(ts)
	// log.Println(ty)
	// ty := fmt.Sprintf("%x", time.Now().Unix())
	// log.Println(ty)
	// strconv.ParseInt(ty, 16, 32)
	// fmt.Println(ty - ts)
	// fmt.Printf("generateRandomArry(10, 5): %v\n", generateRandomArry(10, 45))
	// fmt.Printf("%v,%v", a.edges[0], a.nodes)
	// node := init_Node(2)
	// a := []string{}
	// if a == nil {
	// 	fmt.Println(12)
	// }
	// sort.Sort(StringSlice(a))
	// var x int = 7   // x has dynamic type int and value 7
	// i := x.(string) // i has type int and value 7
	// a := 255 & (^(0 | 0 | 0))
	// a := 13 >> 1
	// fmt.Printf("%b\n", 209)
	// fmt.Printf("%b", 162|1)

	// pos     11010000
	// 		   01110000
	// ~pos+1  01011101
	// isConnect := [][]int{{1, 2, 3}, {5, 6, 7}}
	// for i, v := range isConnect {
	// 	//     rank[i] = 1
	// 	//     parent[i] = isConnect[]
	// 	fmt.Println(v[i])

	// }

	// }
	// fmt.Println(Decimal(8, 3))
	// result := math.Round(math.Pi*1e3) / 1e3
	// fmt.Printf("result: %v\n", result)
	// fmt.Println(math.Pow(float64(8), 2))
	// ans[0][4] = make([]int, 3)
	// ans[0][4] = append(ans[0][4], 22)
	// nums := []int{5, 4, 6}
	// pos & (^pos + 1)
	// a := 52
	// arr := make([]int, (rand.Int63()*len)+1)
	// arr := []int{1, 2, 3}
	// var dp [][]int
	// dp = make([][]int, len(arr)+1)

	// fmt.Printf("dp: %v\n", dp)
	// ch1 := []byte{}
	// fmt.Printf("%b\n", pos-1)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%b\n", ^pos+1)
	// fmt.Printf("%b\n", pos&(^pos+1))

	// 00001000
	// fmt.Println((n >> 31 & 1) ^ 1)

	// link = &Link{3, nil}
	// stack = append(stack, link)
	// link = &Link{5, nil}
	// stack = append(stack, link)
	// a := ans[]
	// ans[0] =
	// fmt.Println(ans)
	// len(ans[0])-1
	// if tmp[p] != tmp[q] {
	// 	cnt = 0
	// 	continue
	//

	// nums = make(map[int]int, 5)
	// // us.parent = make(map[int]int)
	// // for i := 0; i < 4; i++ {
	// // 	nums[i] = 1
	// // }
	// _, ok := nums[0]
	// if !ok {
	// 	fmt.
	// }
	// limit := 1<<4 - 1
	// fmt.Printf("%b", limit)
	// num := make()
	// fmt.Println(time.Weekday())
	// f := []int{1, 2, 3, 4, 5}
	// res := [][]int{}
	// ans[1] = "ss"
	// res = append(res, a)
	// res = append(res, f)
	// fmt.Printf("%T", n[0])
	// for _, v := range a {
	// 	ans += v
	// }
	// fmt.Println(ans)
	// copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中

	// copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
}

func knightProbability2(n, k, row, column int) float64 {
	// 初始化表
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
	}
	// for i := 0; i < row; i++ {
	// 	for j := 0; j < column; j++ {
	// 		dp[i][j] = make([]int, k+1)
	// 	}
	// }
	// for i := 0; i < row; i++ {
	// 	for j := 0; j < column; j++ {
	// 		for q := 0; q < k; q++ {
	// 			dp[i][j][q] = -1
	// 		}
	// 	}
	// }
	ans := process1(n, k, row, column, dp)
	result := math.Round(float64(ans)/math.Pow(float64(8), float64(k))*1e5) / 1e5
	return result

}

func process1(n, k, row, column int, dp [][][]int) int {
	if dp[row][column][k] != 0 {
		return dp[row][column][k]
	}
	if row < 0 || row >= n || column < 0 || column >= n {
		dp[row][column][k] += 0
		return 0
	}
	if k == 0 {
		dp[row][column][k] += 1
		return 1
	}
	live := process1(n, k-1, row+1, column+2, dp)
	live += process1(n, k-1, row-1, column+2, dp)
	live += process1(n, k-1, row+2, column+1, dp)
	live += process1(n, k-1, row-2, column+1, dp)
	live += process1(n, k-1, row+2, column-1, dp)
	live += process1(n, k-1, row-2, column-1, dp)
	live += process1(n, k-1, row+1, column-2, dp)
	live += process1(n, k-1, row-1, column-2, dp)
	return live

}
func Decimal(a, b int) float64 {
	num := math.Round(float64(a) / float64(b))
	// num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(a)/float64(b)), 64)
	return num
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func TopVal(head *Link) int {
	return head.index
}
func f1() *int {
	i := 1
	return &i
}

// func hello(name string) error {
// 	if len(name) == 0 {
// 		return errors.New("error: name is null")
// 	}
// 	// fmt.Println("Hello,", name)
// 	return nil
// }
func test(slice *[]string) {
	(*slice) = append((*slice), "a")
	(*slice)[0] = "b"
	(*slice)[1] = "b"
	fmt.Print(*slice)
}

// func test(hash map[int]int) {
// 	hash[1] = 2
// 	hash[2] = 2
// 	fmt.Println(&hash)

// }
