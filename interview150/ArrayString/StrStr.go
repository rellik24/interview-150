package interview150

import (
	"fmt"
	"strings"
)

// Find the Index of the First Occurrence in a String
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/submissions/1489555554/?envType=study-plan-v2&envId=top-interview-150
func StrStr() {
	fmt.Println(strStr("a", "a"))
}

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	for i := 0; i < l1; i++ {
		if i+l2 <= l1 && strings.HasPrefix(haystack[i:i+l2], needle) {
			return i
		}
	}
	return -1
}
