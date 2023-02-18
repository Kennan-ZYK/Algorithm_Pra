package chapter4

// 普里姆算法
/*
随机选个结点，然后把该结点的边全部
用小根堆把该结点的边存储，然后把边拿出来，看 edge.to是不是新结点，不是新跳过。
若是加入集合中，这个边就要它，然后把 edge.to的所有边存到小根堆里，
依次反复执行以上操作，直到全部点弄完
*/
type Graph struct {
	nodes map[int]*Node
	edges []*Edge
}
type Node struct {
	val   int
	in    int
	out   int
	nexts []*Node
	edges []*Edge
}

type Edge struct {
	weight int
	from   *Node
	to     *Node
}

func heapSmall(arr []int, index int) []int {
	for arr[index] < arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
	return arr
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// func primMST(graph Graph) {

// }
