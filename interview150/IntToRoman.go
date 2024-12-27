package interview150

import (
	"fmt"
)

// Integer to Roman
// https://leetcode.com/problems/integer-to-roman/description/?envType=study-plan-v2&envId=top-interview-150
func IntToRoman() {
	fmt.Println(intToRoman(3749))
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			result += symbols[i]
		}
	}

	return result
}
