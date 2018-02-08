package bitwise

import "testing"

func TestLeftShift(t *testing.T) {
	actual := LeftShift(1)
	expected := 2
	if  actual != expected {
		t.Errorf("expected: %d, actual: %d", actual, expected)
	}

	actual = LeftShift(2)
	expected = 4
	if  actual != expected {
		t.Errorf("expected: %d, actual: %d", actual, expected)
	}

	actual = LeftShift(4)
	expected = 8
	if  actual != expected {
		t.Errorf("expected: %d, actual: %d", actual, expected)
	}
	// A PETTERN EMERGES
}

func TestRightShift(t *testing.T) {
	actual := RightShift(4)
	expected := 2
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", actual, expected)
	}

	actual = RightShift(3)
	expected = 1
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", actual, expected)
	}
}
