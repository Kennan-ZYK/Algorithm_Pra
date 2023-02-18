package chapter4

import "sort"

/*
 每个边按权重由小到大排序，从最小边开始判断，
 若有此边，不形成环--记录，形成环--丢弃。
 如何判断环?
 用集合方式，如 A,B,C,D,E有5条点，分别为一个集合 {A}{B}{C}{D}{E}
 按排序后的边开始，当两个点之间有边时，判断两个点有没在同一个集合里，没有就把两个集合合并
 有则说明形成环
*/

type Vertex int

// Edge describes the edge of a weighted graph
type Edge struct {
	Start  Vertex
	End    Vertex
	Weight int
}

type DisjointSetUnionElement struct {
	Parent Vertex
	Rank   int
}

type DisjointSetUnion []DisjointSetUnionElement

func NewDSU(n int) *DisjointSetUnion {

	dsu := DisjointSetUnion(make([]DisjointSetUnionElement, n))
	return &dsu
}

// MakeSet will create a set in the DSU for the given node
func (dsu DisjointSetUnion) MakeSet(node Vertex) {

	dsu[node].Parent = node
	dsu[node].Rank = 0
}

func (dsu DisjointSetUnion) FindSetRepresentative(node Vertex) Vertex {

	if node == dsu[node].Parent {
		return node
	}

	dsu[node].Parent = dsu.FindSetRepresentative(dsu[node].Parent)
	return dsu[node].Parent
}

func (dsu DisjointSetUnion) UnionSets(firstNode Vertex, secondNode Vertex) {

	firstNode = dsu.FindSetRepresentative(firstNode)
	secondNode = dsu.FindSetRepresentative(secondNode)

	if firstNode != secondNode {

		if dsu[firstNode].Rank < dsu[secondNode].Rank {
			firstNode, secondNode = secondNode, firstNode
		}
		dsu[secondNode].Parent = firstNode

		if dsu[firstNode].Rank == dsu[secondNode].Rank {
			dsu[firstNode].Rank++
		}
	}
}

func KruskalMST(n int, edges []Edge) ([]Edge, int) {

	var mst []Edge // The resultant minimum spanning tree
	var cost int = 0

	dsu := NewDSU(n)

	for i := 0; i < n; i++ {
		dsu.MakeSet(Vertex(i))
	}

	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	for _, edge := range edges {

		if dsu.FindSetRepresentative(edge.Start) != dsu.FindSetRepresentative(edge.End) {

			mst = append(mst, edge)
			cost += edge.Weight
			dsu.UnionSets(edge.Start, edge.End)
		}
	}

	return mst, cost
}
