package interview150

import (
	"fmt"
)

// Number of 1 Bits
// https://leetcode.com/problems/number-of-1-bits/description/?envType=study-plan-v2&envId=top-interview-150
func HammingWeight() {
	fmt.Println(hammingWeight(2147483645))
}

func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}
	return count
}
