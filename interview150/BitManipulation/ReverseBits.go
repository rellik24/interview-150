package interview150

import (
	"fmt"
)

// Reverse Bits
// https://leetcode.com/problems/reverse-bits/?envType=study-plan-v2&envId=top-interview-150
func ReverseBits() {
	fmt.Println(reverseBits(43261596)) // 00000010100101000001111010011100
}

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = (res << 1) | (num & 1)
		num >>= 1
	}
	return res
}
