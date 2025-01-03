package interview150

import (
	"fmt"
	"strings"
)

// Word Pattern
// https://leetcode.com/problems/word-pattern/description/?envType=study-plan-v2&envId=top-interview-150
func WordPattern() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}

func wordPattern(pattern string, s string) bool {
	s1 := strings.Split(s, " ")

	if len(pattern) != len(s1) {
		return false
	}
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	idx1, idx2 := 0, 0
	val1, val2 := 0, 0

	for i := 0; i < len(s1); i++ {
		if v, ok := m1[string(s1[i])]; !ok {
			m1[string(s1[i])] = idx1
			val1 = idx1
			idx1++
		} else {
			val1 = v
		}

		if v, ok := m2[string(pattern[i])]; !ok {
			m2[string(pattern[i])] = idx2
			val2 = idx2
			idx2++
		} else {
			val2 = v
		}
		if val1 != val2 {
			return false
		}
	}
	return true
}
