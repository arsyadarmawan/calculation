package calculation

func Pow(arg int, arg0 int) int {
	if arg0 < 0 {
		return 0
	}
	res := 1
	for arg0 != 0 {
		if (arg0 & 1) == 1 {
			res *= arg
		}
		arg0 >>= 1
		arg *= arg
	}

	if res < 0 {
		panic("Integer overflow, use 'math' library for more big calculations")
	}
	return res
}
