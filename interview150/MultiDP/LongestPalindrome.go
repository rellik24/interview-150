package interview150

import "fmt"

func LongestPalindrome() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	n := len(s)
	if n < 1 {
		return ""
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLen := 0, 1

	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if maxLen < j-i+1 {
					start = i
					maxLen = j - i + 1
				}
			}
		}
	}
	return s[start : start+maxLen]
}
