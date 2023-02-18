package chapter4

type Graph struct {
	nodes map[int]*Node //编号及结点
	edges []*Edge       //所属边
}

type Node struct {
	val   int     // 数据项
	in    int     // 入度
	out   int     // 出度
	nexts []*Node // 邻接结点
	edges []*Edge // 所属边 （有向图：出度边，无向图：有就是）
}

type Edge struct {
	weight int   //权重
	from   *Node //出发点
	to     *Node //终点
}

func init_Graph() *Graph {
	return &Graph{map[int]*Node{}, []*Edge{}}

}

func init_Node(value int) *Node {
	node := &Node{}
	node.val = value
	node.in = 0
	node.out = 0
	node.nexts = []*Node{}
	node.edges = []*Edge{}
	return node
}

func init_Edge(weight int, from *Node, to *Node) *Edge {
	return &Edge{weight, from, to}
}
