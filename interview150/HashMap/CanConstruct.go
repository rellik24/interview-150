package interview150

import (
	"fmt"
	"strings"
)

// Ransom Note
// https://leetcode.com/problems/ransom-note/description/?envType=study-plan-v2&envId=top-interview-150
func CanConstruct() {
	fmt.Println(canConstruct("ab", "b"))
}

func canConstruct(ransomNote string, magazine string) bool {
	for _, note := range magazine {
		ransomNote = strings.Replace(ransomNote, string(note), "", 1)
	}
	return len(ransomNote) == 0
}
