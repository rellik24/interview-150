package interview150

import (
	"fmt"
	"regexp"
	"strings"
)

// Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/description/?envType=study-plan-v2&envId=top-interview-150
func IsPalindrome() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	l := len(s) - 1
	for i, v := range s {
		if i == l-i {
			return true
		}
		if v != rune(s[l-i]) {
			return false
		}
	}
	return true
}
