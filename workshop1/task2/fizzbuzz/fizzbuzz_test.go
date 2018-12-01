package fizzbuzz

import "testing"

func testImpl(t *testing.T, count int, expected string) {
	if Get(count) != expected {
		t.Errorf("Result mismatch")
	}
}

func TestEmpty(t *testing.T) {
	testImpl(t, 0, "")
}

func TestNegative(t *testing.T) {
	testImpl(t, -42, "")
}

func TestFive(t *testing.T) {
	testImpl(t, 5, "1, 2, Fizz, 4, Buzz")
}
