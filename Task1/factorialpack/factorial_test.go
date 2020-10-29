package factorialpack

import "testing"

func TestFactorial(t *testing.T) {
	factorialOfZero := Factorial(0)
	if factorialOfZero != 1 {
		t.Errorf("Factorial(0) returned unexpected result: %d\n", factorialOfZero)
	}
	testCaseOne := Factorial(4)
	if testCaseOne != 24 {
		t.Errorf("Factorial(4) returned unexpected result: %d\n", factorialOfZero)
	}
}
