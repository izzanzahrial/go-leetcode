package graph

import "sort"

// https://leetcode.com/problems/accounts-merge/
func accountsMerge(accounts [][]string) [][]string {
	parent := make([]int, len(accounts))

	// initialize parent
	for i := 0; i < len(accounts); i++ {
		parent[i] = i
	}

	// find with path compression
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// union
	union := func(x, y int) {
		px := find(x)
		py := find(y)
		if px != py {
			parent[py] = px
		}
	}

	// map email -> account index
	emailToIndex := make(map[string]int)

	for i, acc := range accounts {
		for j := 1; j < len(acc); j++ {
			email := acc[j]
			if idx, exists := emailToIndex[email]; exists {
				union(i, idx)
			} else {
				emailToIndex[email] = i
			}
		}
	}

	// group emails by root parent
	rootToEmails := make(map[int][]string)

	for email, idx := range emailToIndex {
		root := find(idx)
		rootToEmails[root] = append(rootToEmails[root], email)
	}

	// build result
	result := [][]string{}

	for root, emails := range rootToEmails {
		sort.Strings(emails)
		name := accounts[root][0]
		account := append([]string{name}, emails...)
		result = append(result, account)
	}

	return result
}
