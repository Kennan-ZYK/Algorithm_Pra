package chapter4

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

func topology(graph Graph) []*Node {
	// 记录每个结点的入度
	Inhash := map[*Node]int{}
	// 记录入度为0的结点
	zeroQueue := []*Node{}
	for i := 0; i < len(graph.nodes); i++ {
		Inhash[graph.nodes[i]] = graph.nodes[i].in
		if graph.nodes[i].in == 0 {
			zeroQueue = append(zeroQueue, graph.nodes[i])
		}
	}

	// 拓扑排序
	res := []*Node{}
	for len(zeroQueue) > 0 {
		cur := zeroQueue[0]
		zeroQueue = zeroQueue[1:]
		res = append(res, cur)
		for i := 0; i < len(cur.nexts); i++ {
			// 把cur的nexts的结点入度 - 1
			Inhash[cur.nexts[i]] -= 1
			// 发现入度为0，加入队列
			if cur.nexts[i].in == 0 {
				zeroQueue = append(zeroQueue, cur.nexts[i])
			}
		}
	}
	return res
}
