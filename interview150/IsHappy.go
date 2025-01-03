package interview150

import (
	"fmt"
)

// Happy Number
// https://leetcode.com/problems/happy-number/description/?envType=study-plan-v2&envId=top-interview-150
func IsHappy() {
	fmt.Println(isHappy(7))
}

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true

		tmp := n
		sum := 0
		for tmp > 0 {
			digit := tmp % 10
			sum += digit * digit
			tmp /= 10
		}
		n = sum
	}
	return n == 1
}
