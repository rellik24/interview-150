package interview150

import (
	"fmt"
)

// Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/description/?envType=study-plan-v2&envId=top-interview-150
func IsIsomorphic() {
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	idx1, idx2 := 0, 0
	val1, val2 := 0, 0

	for i := 0; i < len(s); i++ {
		if v, ok := m1[string(s[i])]; !ok {
			m1[string(s[i])] = idx1
			val1 = idx1
			idx1++
		} else {
			val1 = v
		}

		if v, ok := m2[string(t[i])]; !ok {
			m2[string(t[i])] = idx2
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
