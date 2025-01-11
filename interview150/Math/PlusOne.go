package interview150

import "fmt"

func PlusOne() {
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	f := false
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if f {
			digits[i]++
			f = false
		}
		if digits[i] >= 10 {
			f = true
			digits[i] %= 10
		}
	}
	if f {
		digits = append([]int{1}, digits...)
	}
	return digits
}
