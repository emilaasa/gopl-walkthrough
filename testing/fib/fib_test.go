package fib

import "testing"

var fibTests = []struct {
	n        int
	expected int
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFib(t *testing.T) {
	for _, ft := range fibTests {
		actual := Fib(ft.n)
		if actual != ft.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", ft.n, ft.expected, actual)
		}
	}
}
