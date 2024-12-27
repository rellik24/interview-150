package interview150

import (
	"fmt"
	"strings"
)

// Word Break
// https://leetcode.com/problems/word-break/description/?envType=study-plan-v2&envId=top-interview-150
func WordBreak() {
	fmt.Println(wordBreak("", []string{"leet", "code"}))
}

func wordBreak(s string, wordDict []string) bool {
	list := make([]bool, len(s)+1)
	list[0] = true
	for i := 1; i <= len(s)+1; i++ {
		for _, word := range wordDict {
			if i+len(word) <= len(s)+1 {
				if list[i-1] && strings.Compare(word, s[i-1:i+len(word)-1]) == 0 {
					list[i+len(word)-1] = true
				}
			}
		}
	}
	return list[len(s)]
}
