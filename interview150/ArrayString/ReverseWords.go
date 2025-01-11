package interview150

import (
	"fmt"
	"strings"
)

// Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string/description/?envType=study-plan-v2&envId=top-interview-150
func ReverseWords() {
	fmt.Println(reverseWords("   the sky is blue"))
}

func reverseWords(s string) string {
	list := strings.Fields(s)
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		tmp := list[i]
		list[i] = list[j]
		list[j] = tmp
	}
	return strings.Join(list, " ")
}
