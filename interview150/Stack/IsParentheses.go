package interview150

import (
	"fmt"
	"strings"
)

// Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/description/?envType=study-plan-v2&envId=top-interview-150
func ValidParentheses() {
	fmt.Println(isValidParentheses("()[]{}"))
}

func isValidParentheses(s string) bool {
	m := make(map[string]int)
	m["{"] = 1
	m["["] = 2
	m["("] = 3
	m[")"] = -3
	m["]"] = -2
	m["}"] = -1

	var arr []int
	for _, v := range strings.Split(s, "") {
		val := m[v]
		if val > 0 {
			arr = append(arr, val)
		} else {
			l := len(arr)
			if l == 0 {
				return false
			} else if arr[l-1]+val == 0 {
				arr = arr[:l-1]
			} else {
				return false
			}
		}
	}
	return len(arr) == 0
}
