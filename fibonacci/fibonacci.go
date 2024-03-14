package fibonacci

func fibonacci1() func() int {
	x, y := 0, 1

	return func() int {
		f := x
		x, y = y, f+y
		return f
	}
}

func fibonacci2(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci2(n-1) + fibonacci2(n-2)
}
