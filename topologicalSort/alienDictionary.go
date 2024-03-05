package main

import "fmt"

func main() {
	AlienDictionary([]string{"abcd", "defg", "ghij", "jklmn", "nopqr"})
}

type graph struct {
	key string
	val []string
}

func AlienDictionary(words []string) string {
	graphs := createGraph(words)

	var noDependecy []*graph
	for _, g := range graphs {
		fmt.Printf("%v: %v\n", g.key, g.val)
		if len(g.val) == 0 {
			noDependecy = append(noDependecy, g)
		}
	}
	for _, g := range noDependecy {
		fmt.Printf("noDependecy %v: %v\n", g.key, g.val)
	}

	var result string
	for len(noDependecy) > 0 {
		curr := noDependecy[0].key
		result = result + noDependecy[0].key
		noDependecy = noDependecy[1:]

		for _, g := range graphs {
			for i, v := range g.val {
				if v == curr {
					g.val = append(g.val[:i], g.val[i+1:]...)
				}
			}

			if len(g.val) == 0 {
				noDependecy = append(noDependecy, g)
			}
		}
	}

	if len(result) == len(graphs) {
		return result
	}

	return ""
}

func createGraph(words []string) []*graph {
	graphs := []*graph{}

	for _, word := range words {
		for _, c := range word {
			isKey := false
			for _, g := range graphs {
				if g.key == string(c) {
					isKey = true
				}
			}

			if !isKey {
				graphs = append(graphs, &graph{key: string(c)})
			}
		}
	}

	// loop through the list of words
	for _, word := range words {

		// get the current graph key
		// search from index 1 because index 0 doesn't have dependencies
		for i := 1; i < len(word); i++ {

			// the key
			key := string(word[i])

			// loop through the graph elements that contain key
			for _, g := range graphs {
				if g.key == key {
					isVal := false

					// if the graph val already contains the previous character(word[i-1]) then skip it
					for _, v := range g.val {
						if v == string(word[i-1]) {
							isVal = true
						}
					}

					// if the graph val does not contain the previous character add the previous character
					if !isVal {
						g.val = append(g.val, string(word[i-1]))
					}
				}
			}
		}
	}

	return graphs
}
