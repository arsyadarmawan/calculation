package calculation

func Max(args ...int) int {
	var max int = args[0]

	for _, score := range args {
		if max < score {
			max = score
		}
	}
	return max
}
