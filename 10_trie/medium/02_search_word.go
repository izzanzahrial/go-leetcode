package trie

type TrieNode2 struct {
	children map[rune]*TrieNode2
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNode2
}

func Constructor2() WordDictionary {
	return WordDictionary{root: &TrieNode2{children: make(map[rune]*TrieNode2)}}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for _, char := range word {
		// if there are no trie with the current character
		// then we created it
		if curr.children[char] == nil {
			curr.children[char] = &TrieNode2{children: make(map[rune]*TrieNode2)}
		}

		// we move the current into the children
		curr = curr.children[char]
	}

	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	curr := this.root
	return this.searchDFS(word, curr)
}

func (this *WordDictionary) searchDFS(word string, node *TrieNode2) bool {
	if len(word) == 0 {
		return node.isEnd
	}

	char := rune(word[0])

	if char == '.' {
		for _, child := range node.children {
			if this.searchDFS(word[1:], child) {
				return true
			}
		}
		return false
	} else {
		if node.children[char] == nil {
			return false
		}
		return this.searchDFS(word[1:], node.children[char])
	}
}
