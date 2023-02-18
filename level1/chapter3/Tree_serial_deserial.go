package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 二叉树序列化和反序列化
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 前序
func SerialByPre(head *Node) string {
	if head == nil {
		return "#_"
	}
	res := strconv.Itoa(head.Val) + "_"
	res = res + SerialByPre(head.Left)
	res = res + SerialByPre(head.Right)
	return res
}

func reconByPre(res string) *Node {
	ans := strings.Split(res, "_")
	Queue := []string{}
	// len(ans)-1 是为了去除最后一个空格
	Queue = append(Queue, ans...)
	Queue = Queue[:len(Queue)-1]
	return ReconPreOrder(&Queue)
}

func ReconPreOrder(Queue *[]string) *Node {
	value := (*Queue)[0]
	(*Queue) = (*Queue)[1:]
	if value == "#" {
		return nil
	}
	ans, _ := strconv.Atoi(value)
	root := &Node{Val: ans}
	root.Left = ReconPreOrder(Queue)
	root.Right = ReconPreOrder(Queue)

	return root
}

func main() {

	head := &Node{Val: 1}
	head.Left = &Node{Val: 2}
	head.Right = &Node{Val: 3}
	// head.Left.Left = &Node{Val: nil}
	// head.Left.Right = &Node{Val: nil}
	head.Right.Left = &Node{Val: 4}
	head.Right.Right = &Node{Val: 5}
	ret := SerialByPre(head)
	fmt.Println(reconByPre(ret))

}
