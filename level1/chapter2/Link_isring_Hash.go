// 哈希表，把经过的点记录下来，当发现有相同的，就确定为进环结点。
//
package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func getIntersectionNode(h1, h2 *LinkNode) *LinkNode {
	vis := map[*LinkNode]bool{}
	for tmp := h1; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := h2; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}
