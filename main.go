package main

import (
	"fmt"
	"math/big"
	"strconv"
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

// reverse takes an integer value, transforms it into a string
// reverses the string and makes an integer out of it.
// So the chain is: 312 -> "312" -> "213" -> 213
func reverse(n *big.Int) *big.Int {
	ns := n.String()
	nsr := strReverse(ns)
	xr := new(big.Int)
	xr, ok := xr.SetString(nsr, 10)
	if !ok {
		panic(fmt.Sprintf("Conversion error to big.Int: %v", nsr))
	}
	return xr
}

func isPalindromNumber(val big.Int) bool {
	h := reverse(&val)
	return val.String() == h.String()
}

func run() {
	N := 10000
	maxDepth := 50

	lychrels := make([]int64, 0)

	for n := 1; n <= N; n++ {
		x := new(big.Int)
		x.SetString(strconv.FormatInt(int64(n), 10), 10)

		for j := 0; j < maxDepth; j++ {
			// Calculate if number + reverse(number) is palindrom
			x = new(big.Int).Add(x, reverse(x))
			isPalindromic := isPalindromNumber(*x)
			if isPalindromic {
				break
			}

			if j == maxDepth-1 {
				lychrels = append(lychrels, int64(n))
			}
		}
	}

	for _, v := range lychrels {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("Found %v lychrel numbers\n", len(lychrels))
}

func main() {
	run()
}
