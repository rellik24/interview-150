package interview150

import "fmt"

func TrailingZeroes() {
	fmt.Println(trailingZeroes(5))
}

func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}
