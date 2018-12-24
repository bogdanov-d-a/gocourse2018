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

func TestNoFizz(t *testing.T) {
	testImpl(t, 2, "1, 2")
}

func TestNoBuzz(t *testing.T) {
	testImpl(t, 4, "1, 2, Fizz, 4")
}

func TestFizzBuzz(t *testing.T) {
	testImpl(t, 16, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16")
}
