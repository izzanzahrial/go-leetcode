package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	var result [][]string

	currRes := make(map[string][]string)

	abc := make(map[string]int)
	for i := 0; i < 26; i++ {

	}

	for _, str := range strs {
		for _, val := range str {
			strVal := string(val)
			abc[strVal] += 1
		}

		var currKey string
		for key, val := range abc {
			if val != 0 {
				currKey += key
			}
		}
		fmt.Println(currKey)

		currRes[currKey] = append(currRes[currKey], str)
	}

	var i int
	for _, str := range currRes {
		for _, val := range str {
			result[i] = append(result[i], val)
		}
		i++
	}

	return result
}
