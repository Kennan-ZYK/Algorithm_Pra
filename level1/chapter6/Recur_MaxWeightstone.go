package main

import "fmt"

/*
题目:给定两个长度都为N的数组weights和values，weights[i]和values[i]分别代表i号物品的重量和价值。
给定一个正数bag，表示一个载重bag的袋子，你装的物品不能超过这个重量。返回你能装下最多的价值是多少?
*/

func process1(weights, values []int, i, alreadyweight, bag int) int {
	if alreadyweight > bag {
		return -values[i-1]
	}
	if i == len(weights) {
		return 0
	}
	return max(
		process1(weights, values, i+1, alreadyweight, bag),
		values[i]+process1(weights, values, i+1, alreadyweight+weights[i], bag))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weights := []int{1, 2, 3, 2}
	values := []int{1, 2, 3, 2}
	bag := 4
	a := process1(weights, values, 0, 0, bag)
	fmt.Println(a)
}
