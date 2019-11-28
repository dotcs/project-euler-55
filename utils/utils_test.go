package utils

import (
	"math/big"
	"testing"
)

func TestStrReverse(t *testing.T) {
	given := strReverse("Foobar")
	expected := "rabooF"
	if given != expected {
		t.Errorf("Expected %v got %v", expected, given)
	}
}

func TestReverse(t *testing.T) {
	n := new(big.Int)
	n.SetString("123", 10)
	given := Reverse(n)
	expected := new(big.Int)
	expected.SetString("321", 10)

	if given.String() != expected.String() {
		t.Errorf("Expected %v got %v", expected, given)
	}
}

func TestIsPalindromNumber_negative(t *testing.T) {
	given := new(big.Int)
	given.SetString("1337", 10)
	expected := false

	if IsPalindromNumber(given) != expected {
		t.Errorf("Expected %v got %v", expected, given)
	}
}

func TestIsPalindromNumber_positive(t *testing.T) {
	given := new(big.Int)
	given.SetString("4334", 10)
	expected := true

	if IsPalindromNumber(given) != expected {
		t.Errorf("Expected %v got %v", expected, given)
	}
}

func TestIsLychrel_false(t *testing.T) {
	given := IsLychrel(int64(47), 1)
	expected := false

	if given != expected {
		t.Errorf("Expected %v got %v", expected, given)
	}
}

func TestIsLychrel_false2(t *testing.T) {
	given := IsLychrel(int64(349), 3)
	expected := false

	if given != expected {
		t.Errorf("Expected %v got %v", expected, given)
	}
}

func TestIsLychrel_true(t *testing.T) {
	given := IsLychrel(int64(196), 50)
	expected := true

	if given != expected {
		t.Errorf("Expected %v got %v", expected, given)
	}
}
