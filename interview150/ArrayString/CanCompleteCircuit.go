package interview150

import "fmt"

// Gas Station
// https://leetcode.com/problems/gas-station/description/?envType=study-plan-v2&envId=top-interview-150
func CanCompleteCircuit() {
	fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost, tank, start := 0, 0, 0, 0

	for i := range gas {
		totalCost += cost[i]
		totalGas += gas[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}
	return start
}
