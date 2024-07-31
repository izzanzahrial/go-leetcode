package array

import (
	"log"
	"strings"
)

// https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75
func MergeAlternately(word1 string, word2 string) string {
	var result string

	for len(word1) != 0 && len(word2) != 0 {
		result = result + string(word1[0])
		result = result + string(word2[0])

		word1 = word1[1:]
		word2 = word2[1:]
	}

	for len(word1) != 0 {
		result = result + string(word1[0])
		word1 = word1[1:]
	}

	for len(word2) != 0 {
		result = result + string(word2[0])
		word2 = word2[1:]
	}

	return result
}

func mergeAlternately(word1 string, word2 string) string {
	var result strings.Builder

	for len(word1) != 0 && len(word2) != 0 {
		err := result.WriteByte(word1[0])
		if err != nil {
			log.Fatal(err)
		}

		err = result.WriteByte(word2[0])
		if err != nil {
			log.Fatal(err)
		}

		word1 = word1[1:]
		word2 = word2[1:]
	}

	for len(word1) != 0 {
		err := result.WriteByte(word1[0])
		if err != nil {
			log.Fatal(err)
		}

		word1 = word1[1:]
	}

	for len(word2) != 0 {
		err := result.WriteByte(word2[0])
		if err != nil {
			log.Fatal(err)
		}

		word2 = word2[1:]
	}

	return result.String()
}
