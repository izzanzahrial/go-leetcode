package pointer

import (
	"log"
	"strings"
)

// https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-
func mergeAlternately3(word1 string, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return word2
	}

	if len2 == 0 {
		return word1
	}

	var result strings.Builder
	word1Idx, word2Idx := 0, 0
	for word1Idx < len1 && word2Idx < len2 {
		result.WriteByte(word1[word1Idx])
		result.WriteByte(word2[word2Idx])
		word1Idx++
		word2Idx++
	}

	if word1Idx < len1 {
		result.WriteString(word1[word1Idx:len1])
	}

	if word2Idx < len2 {
		result.WriteString(word2[word2Idx:len2])
	}

	return result.String()
}

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
