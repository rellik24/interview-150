package interview150

import "fmt"

// Valid Anagram
// https://leetcode.com/problems/valid-anagram/description/?envType=study-plan-v2&envId=top-interview-150
func IsAnagram() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	list := make([]int, 26)
	for i := 0; i < len(s); i++ {
		list[s[i]-'a']++
		list[t[i]-'a']--
	}

	for _, v := range list {
		if v != 0 {
			return false
		}
	}
	return true
}
