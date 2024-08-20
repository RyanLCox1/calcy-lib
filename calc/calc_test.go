package calc

import "testing"

func TestAddition(t *testing.T) {
	a := Addition{}
	assertEqual(t, 3, a.Calculate(2, 1))
	assertEqual(t, 13, a.Calculate(3, 10))
	assertEqual(t, -1, a.Calculate(-2, 1))
	assertEqual(t, 0, a.Calculate(0, 0))
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Helper()
		t.Errorf("%v != %v", a, b)
	}
}
