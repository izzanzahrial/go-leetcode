package trie

// https://leetcode.com/problems/implement-trie-prefix-tree/
type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, char := range word {
		// if there are no trie with the current character
		// then we created it
		if curr.children[char] == nil {
			curr.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}

		// we move the current into the children
		curr = curr.children[char]
	}

	curr.endOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, c := range word {
		if curr.children[c] == nil {
			return false
		}

		curr = curr.children[c]
	}

	return curr.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, c := range prefix {
		if curr.children[c] == nil {
			return false
		}

		curr = curr.children[c]
	}

	return true
}
