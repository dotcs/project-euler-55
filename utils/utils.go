package utils

import (
	"math/big"
)

// strReverse takes a string and reverses this string, e.g.
// "foobar" -> "raboof"
func strReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Reverse takes an integer value, transforms it into a string
// reverses the string and makes an integer out of it.
// So the chain is: 312 -> "312" -> "213" -> 213
func Reverse(n *big.Int) *big.Int {
	ns := n.String()
	nsr := strReverse(ns)
	xr := new(big.Int)
	xr.SetString(nsr, 10)
	return xr
}

// IsPalindromNumber returns true if the given number is the same
// when being read from left to right and from right to left.
// E.g. 121 is a palindrom number whereas 143 is not.
func IsPalindromNumber(n *big.Int) bool {
	h:=n.String()
	l:=len(h)
	result := true
	for i := 0; i < l/2; i++ {
		if h[i] != h[l-i-1] {
			result=false
			break
		}
	}
	return result
}
