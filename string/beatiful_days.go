package string

import (
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem?isFullScreen=true
func BeautifulDays(i int32, j int32, k int32) int32 {
	var slayDays int32
	for l := i; l <= j; l++ {
		revStrNum := reverseNum(l)
		revNum, _ := strconv.Atoi(revStrNum)
		dayValue := l - int32(revNum)
		if dayValue%k == 0 {
			slayDays++
		}
	}

	return slayDays
}

func reverseNum(num int32) string {
	var str strings.Builder
	strNum := strconv.FormatInt(int64(num), 10)
	for i := len(strNum) - 1; i >= 0; i-- {
		str.WriteString(string(strNum[i]))
	}
	return str.String()
}
