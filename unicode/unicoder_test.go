package main

import "testing"

var unicodeTests = []struct {
	s        string
	expected rune
}{
	{"a", 97},
	{"b", 98},
	{"ä", 228},
}

func TestUnicoder(t *testing.T) {
	for _, ut := range unicodeTests {
		actual := Unicoder(ut.s)
		if actual != ut.expected {
			t.Errorf("Unicoder(\"ut.s\") should return %d but instead returned %d", ut.expected, actual)
		}
	}
}
