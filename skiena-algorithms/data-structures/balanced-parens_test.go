package datastructures

import "testing"

var parenTests = []struct {
	s        string
	expected bool
}{
	{"(())", true},
	{"())", false},
}

func TestIsBalancedParen(t *testing.T) {
	for _, p := range parenTests {
		actual := IsBalancedParen(p.s)
		if actual != p.expected {
			t.Errorf("IsBalancedParen(\"%s\") should be: %t, was:  %v", p.s, p.expected, actual)
		}
	}
}
