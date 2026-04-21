package trie

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

func SearchWithinTheBoard(curr string, row, col int, board [][]byte, node *TrieNode, result map[string]bool) {
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || board[row][col] == '#' {
		return
	}

	char := rune(board[row][col])
	next := node.children[char]
	if next == nil {
		return
	}

	curr += string(char)
	if next.endOfWord {
		result[curr] = true
	}

	// mark visited
	temp := board[row][col]
	board[row][col] = '#'

	// dfs explore all directions
	SearchWithinTheBoard(curr, row-1, col, board, next, result)
	SearchWithinTheBoard(curr, row+1, col, board, next, result)
	SearchWithinTheBoard(curr, row, col-1, board, next, result)
	SearchWithinTheBoard(curr, row, col+1, board, next, result)

	// backtrack
	board[row][col] = temp
}

func findWords(board [][]byte, words []string) []string {
	// create the trie and insert the words
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	result := make(map[string]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			SearchWithinTheBoard("", i, j, board, trie.root, result)
		}
	}

	var result2 []string
	for k, v := range result {
		if v {
			result2 = append(result2, k)
		}
	}

	return result2
}
