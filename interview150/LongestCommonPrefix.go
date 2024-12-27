package interview150

import (
	"fmt"
	"strings"
)

// Longest Common Prefix
// https://leetcode.com/problems/longest-common-prefix/submissions/1489538436/?envType=study-plan-v2&envId=top-interview-150
func LongestCommonPrefix() {
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	i := 1
	for i < len(strs) {
		if !strings.HasPrefix(strs[i], prefix) {
			if len(prefix) >= 1 {
				prefix = prefix[:len(prefix)-1]
			}
			continue
		} else {
			i++
		}
	}
	return prefix
}
