package chapter3

type TrieNode struct {
	pass  int
	end   int
	nexts map[rune]*TrieNode
}

func Constructor() TrieNode {
	root := new(TrieNode)
	root.nexts = make(map[rune]*TrieNode)
	return *root
}
func (this *TrieNode) Insert(word string) {
	if len(word) == 0 {
		return
	}
	node := this
	node.pass++
	for _, ch := range word {
		if node.nexts[ch] == nil {
			newnode := new(TrieNode)
			newnode.nexts = make(map[rune]*TrieNode)
			node.nexts[ch] = newnode
		}
		node = node.nexts[ch]
		node.pass++
	}
	node.end++

}
func (t *TrieNode) Search(word string) int {
	node := t
	for _, ch := range word {
		// word是新词从未统计过
		if node.nexts[ch] == nil {
			return 0
		}
		node = node.nexts[ch]
	}
	return node.end
}

func (t *TrieNode) Delete(word string) {
	if t.Search(word) != 0 {
		node := t
		for _, ch := range word {
			if node.nexts[ch].pass-1 == 0 {
				node.nexts[ch].pass--
				node.nexts[ch] = nil
				return
			} else {
				node.nexts[ch].pass--
			}
			node = node.nexts[ch]
		}
		node.end--
	}

}

// -----------------leetcode-----------------------//

/** Initialize your data structure here. */
// func Constructor() TrieNode {
// 	root := new(TrieNode)
// 	root.nexts = make(map[rune]*TrieNode)
// 	return *root
// }

/** Returns if the word is in the trie. */
// func (this *TrieNode) Search(word string) bool {
//     node := this
//     for _ ,ch := range word{
//         if node.nexts[ch] == nil{
//             return false
//         }
//         node = node.nexts[ch]
//     }
//     if node.end < 1{
//         return false
//     }
//     return true
// }

/** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {
// 	node := this
// 	for _, ch := range prefix {
// 		if node.nexts[ch] == nil {
// 			return false
// 		}
// 		node = node.nexts[ch]
// 	}
// 	return true
// }
