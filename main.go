package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{0, 2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	l := len(nums)
	list := make([]bool, l)
	for i, v := range nums {
		for j := 1; j <= v; j++ {
			if i+j <= l-1 && list[i] {
				list[i+j] = true
			} else {
				break
			}
		}
	}
	fmt.Println(list)
	return list[l-1]
}
