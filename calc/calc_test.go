package calc

import "testing"

func TestAddition(t *testing.T) {
	a := Addition{}
	assertEqual(t, 3, a.Calculate(2, 1))
	assertEqual(t, 13, a.Calculate(3, 10))
	assertEqual(t, -1, a.Calculate(-2, 1))
	assertEqual(t, 0, a.Calculate(0, 0))
}

func TestSubtraction(t *testing.T) {
	s := Subtraction{}
	assertEqual(t, 1, s.Calculate(2, 1))
	assertEqual(t, -7, s.Calculate(3, 10))
	assertEqual(t, -3, s.Calculate(-2, 1))
	assertEqual(t, 0, s.Calculate(0, 0))
}

func TestMultiplication(t *testing.T) {
	m := Multiplication{}
	assertEqual(t, 2, m.Calculate(2, 1))
	assertEqual(t, 30, m.Calculate(3, 10))
	assertEqual(t, -2, m.Calculate(-2, 1))
	assertEqual(t, 0, m.Calculate(0, 0))
}

func assertEqual(t *testing.T, a, b any) {
	if a != b {
		t.Helper()
		t.Errorf("%v != %v", a, b)
	}
}
