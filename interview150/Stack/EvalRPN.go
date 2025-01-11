package interview150

import (
	"fmt"
	"strconv"
)

// Evaluate Reverse Polish Notation
// https://leetcode.com/problems/evaluate-reverse-polish-notation/?envType=study-plan-v2&envId=top-interview-150
// FILO ，如果是數字就塞入Stack，如果是運算符號就取前兩個運算，這題的重點只有看懂題目表示法
func EvalRPN() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	var stack []int

	for _, v := range tokens {
		if val, err := strconv.Atoi(v); err != nil {
			last := len(stack) - 1
			switch v {
			case "+":
				stack[last-1] = stack[last-1] + stack[last]
			case "-":
				stack[last-1] = stack[last-1] - stack[last]
			case "*":
				stack[last-1] = stack[last-1] * stack[last]
			case "/":
				stack[last-1] = stack[last-1] / stack[last]
			}
			stack = stack[:last]
		} else {
			stack = append(stack, val)
		}
	}
	return stack[0]
}
