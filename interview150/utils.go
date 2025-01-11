package interview150

func GetMin(x, y int) int {
	return getMin(x, y)
}

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func GetMax(x, y int) int {
	return getMax(x, y)
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
