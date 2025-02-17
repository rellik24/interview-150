package interview150

import (
	"fmt"
	"strconv"
	"strings"
)

// Add Binary
// https://leetcode.com/problems/add-binary/?envType=study-plan-v2&envId=top-interview-150
func AddBinary() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	if lb > la {
		a, b, la, lb = b, a, lb, la
	}

	b = strings.Repeat("0", la-lb) + b

	arr1 := strings.Split(a, "")

	var carry int
	for i := len(arr1) - 1; i >= 0; i-- {
		v1, _ := strconv.Atoi(arr1[i])
		v2, _ := strconv.Atoi(string(b[i]))
		sum := v1 + v2 + carry
		arr1[i], carry = strconv.Itoa(sum%2), sum/2

	}
	if carry == 1 {
		arr1 = append([]string{"1"}, arr1...)
	}
	return strings.Join(arr1, "")
}
