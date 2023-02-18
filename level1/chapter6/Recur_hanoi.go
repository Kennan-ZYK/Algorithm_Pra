package chapter6

import (
	"fmt"
)

func hannoi(n int) {
	if n < 0 {
		return
	}
	process(n, "左", "中", "右")
}

func process(n int, start, other, end string) {
	if n == 1 {
		fmt.Printf("%s -> %s\n", start, end)
		return
	}
	// 把1..n-1放到other
	process(n-1, start, end, other)
	// 再把n放到end
	fmt.Printf("%s -> %s\n", start, end)
	// 把other在放到end
	process(n-1, other, start, end)

}

func main() {
	// var n int
	// fmt.Scanf("%d", &n)
	hannoi(3)
}
