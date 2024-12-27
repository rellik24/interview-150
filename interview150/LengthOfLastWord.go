package interview150

import "fmt"

// Length of Last Word
// https://leetcode.com/problems/length-of-last-word/description/?envType=study-plan-v2&envId=top-interview-150
func LengthOfLastWord() {
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if count == 0 {
				continue
			}
			break
		} else {
			count++
		}
	}
	return count
}
