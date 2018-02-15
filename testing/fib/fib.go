package fib

func Fib(n int) int {
	if n < 3 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
