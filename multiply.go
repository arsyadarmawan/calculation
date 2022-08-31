package calculation

func Multiply(val1 int, val2 int) int {
	result := val1 * val2

	if result < 0 {
		panic("result less indicate negative vallue")
	}

	if val1 < 0 || val2 < 0 {
		panic("input less than zero value")
	}

	return result
}
