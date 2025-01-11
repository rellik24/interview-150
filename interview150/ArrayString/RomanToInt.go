package interview150

import (
	"fmt"
	"strings"
)

// Roman to Integer
// https://leetcode.com/problems/roman-to-integer/description/?envType=study-plan-v2&envId=top-interview-150
func RomanToInt() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := 0

	for i, v := range symbols {
		for strings.HasPrefix(s, v) {
			s = strings.Replace(s, v, "", 1)
			result += values[i]
		}
	}

	return result
}
