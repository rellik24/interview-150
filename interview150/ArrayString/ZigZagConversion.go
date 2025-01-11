package interview150

import (
	"fmt"
	"strings"
)

// Zigzag Conversion
// https://leetcode.com/problems/zigzag-conversion/description/?envType=study-plan-v2&envId=top-interview-150
func ZigZagConversion() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	list := make([]string, numRows)
	direction, count := 1, 0

	for i := 0; i < len(s); i++ {
		list[count] += string(s[i])
		if count == 0 {
			direction = 1
		} else if count == numRows-1 {
			direction *= -1
		}
		count += direction
	}
	return strings.Join(list, "")
}
