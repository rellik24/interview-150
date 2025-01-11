package interview150

import "fmt"

// Is Subsequence
// https://leetcode.com/problems/is-subsequence/?envType=study-plan-v2&envId=top-interview-150
func IsSubsequence() {
	fmt.Println(isSubsequence("", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	idx := 0
	l := len(s)
	for _, v := range t {
		if idx == l {
			return true
		}
		if v == rune(s[idx]) {
			idx++
		}
	}
	return idx == l

}
