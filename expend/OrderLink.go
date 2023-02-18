package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Key   uint64
	Value string
}

func SeachNodeList(arr []Node, x uint64) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i].Key >= x
	})
}

func InsertOrderedNodeList(arr []Node, x Node) []Node {

	i := SeachNodeList(arr, x.Key)

	arr = append(arr, Node{})
	copy(arr[i+1:], arr[i:])
	arr[i] = x

	return arr
}

func main() {
	var arr []Node

	arr = InsertOrderedNodeList(arr, Node{Key: 12, Value: "ha"})
	arr = InsertOrderedNodeList(arr, Node{Key: 1432, Value: "ta"})
	arr = InsertOrderedNodeList(arr, Node{Key: 112, Value: "see"})
	arr = InsertOrderedNodeList(arr, Node{Key: 1452, Value: "sight"})
	arr = InsertOrderedNodeList(arr, Node{Key: 13332, Value: "peer"})
	arr = InsertOrderedNodeList(arr, Node{Key: 12, Value: "one"})
	arr = InsertOrderedNodeList(arr, Node{Key: 2, Value: "kiss"})

	for _, node := range arr {
		fmt.Println(node)
	}

	// {2 kiss}
	// {12 one}
	// {12 ha}
	// {112 see}
	// {1432 ta}
	// {1452 sight}
	// {13332 peer}
}
