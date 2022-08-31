package calculation

func Sum(args ...int) int {
	var val int = args[0]
	for _, score := range args {
		val += score
	}
	return val
}
