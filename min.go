package calculation

func Min(args ...int) int {
	res := MinCalculation(args)
	return res
}

func MinCalculation(a []int) int {
	var min int = a[0]
	for _, val := range a {
		if min > val {
			min = val
		}
	}
	return min
}
