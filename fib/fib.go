package fib

func Fib(a int) int {
	if a <= 1 {
		return a
	}
	return Fib(a-1) + Fib(a-2)
}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second,
			first+second
		return ret
	}
}
