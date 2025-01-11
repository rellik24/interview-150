package interview150

import (
	"fmt"
)

// Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-interview-150
func LengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func lengthOfLongestSubstring(s string) int {
	left, right, max := 0, 0, 0
	m := make(map[string]int)
	for right < len(s) {
		str := string(s[right])
		if _, exist := m[str]; exist { // 存在，左邊收縮
			if max < right-left {
				max = right - left
			}
			for left <= right {
				if string(s[left]) != str {
					delete(m, string(s[left]))
					left++
				} else {
					delete(m, string(s[left]))
					left++
					break
				}
			}
		} else { // 不存在，右邊擴展
			m[str] = 1
			right++
		}
	}
	if max < right-left {
		max = right - left
	}
	return max
}
