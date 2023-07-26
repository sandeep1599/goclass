package fibonacci

func FibonacciSeries(n int) []int {
	if n <= 0 {
		return []int{}
	}

	fib := []int{0, 1}

	for i := 2; i < n; i++ {
		newNum := fib[i-1] + fib[i-2]
		fib = append(fib, newNum)
	}

	return fib
}
