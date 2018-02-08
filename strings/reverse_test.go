package strings

import "testing"

func TestReverseString(t *testing.T) {
	exp := "evol"
	act := ReverseString("love")
	report(t, exp, act)
}

func TestShiftLetters(t *testing.T) {
	exp := "bcd"
	act := ShiftLetters("abc")
	report(t, exp, act)

	exp = "cde"
	act = ShiftLetters("bcd")
	report(t, exp, act)

}

func report(t *testing.T, exp, act string) {
	if act != exp {
		t.Errorf("Yeepers! actual: %s was not equal to the expected: %s", act, exp)
	}
}
