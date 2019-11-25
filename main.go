package main

import (
	"fmt"
	"strconv"
)

type result struct {
	palindromic bool
	atIteration int
	origin      int
}

type lychrel bool

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
func reverse(n int64) int64 {
	ns := strconv.FormatInt(n, 10)
	nsr := strReverse(ns)
	xr, err := strconv.ParseInt(nsr, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Conversion error: %v", err))
	}
	return xr
}

func isPalindromNumber(val int64) bool {
	return val == reverse(val)
}

func main() {
	N := 5000
	maxDepth := 20

	cache := make(map[int64]result)
	lychrels := make([]int64, 0)

	for n := 1; n <= N; n++ {
		x := int64(n)

		for j := 0; j < maxDepth; j++ {
			r, ok := cache[x]
			if ok && r.atIteration <= j {
				// This number has been calculated already
				fmt.Printf("Skip number %v\n", x)
				break
			}

			// Calculate if number + reverse(number) is palindrom
			x = x + reverse(x)
			isPalindromic := isPalindromNumber(int64(x))
			if isPalindromic {
				break
			}

			// cache[int64(x)] = result{palindromic: isPalindromic, atIteration: j, origin: n}

			if j == maxDepth-1 {
				lychrels = append(lychrels, int64(n))
			}
		}
	}

	// for k, v := range cache {
	// 	if v.palindromic {
	// 		fmt.Printf("%v: %v\n", k, v)
	// 	}
	// }

	for _, v := range lychrels {
		fmt.Printf("%v\n", v)
	}
}
