package chapter3

// 前缀树

type TrieNode struct {
	pass int
	end  int
	// nexts[0] ,nexts[1]...nexts[25] 下标代表a,b....z
	nexts [26]*TrieNode
}

func (t *TrieNode) Insert(word string) {
	if len(word) == 0 {
		return
	}
	node := t
	node.pass++
	for _, ch := range word {
		index := ch - 'a'
		// 如果没有记录过，就创建新结点
		if node.nexts[index] == nil {
			node.nexts[index] = new(TrieNode)
		}
		// node就往下走
		node = node.nexts[index]
		node.pass++
	}
	node.end++
}

// 查词频
func (t *TrieNode) Search(word string) int {
	if len(word) == 0 {
		return 0
	}
	node := t
	for _, ch := range word {
		index := ch - 'a'
		// word是新词从未统计过
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.end
}

func (t *TrieNode) Delete(word string) {
	if t.Search(word) != 0 {
		node := t
		for _, ch := range word {
			index := ch - 'a'
			// 若当前结点为最后一个字母，当前结点为nil
			if node.nexts[index].pass-1 == 0 {
				node.nexts[index].pass--
				node.nexts[index] = nil
				return
			} else {
				node.nexts[index].pass--
			}
			node = node.nexts[index]
		}
		node.end--
	}
}
