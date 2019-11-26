package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/dotcs/project-euler-55/utils"
)

func run(N, maxDepth int64) {
	lychrels := make([]int64, 0)

	var n, j int64

	for n = 1; n <= N; n++ {
		x := new(big.Int)
		x.SetString(strconv.FormatInt(n, 10), 10)

		for j = 0; j < maxDepth; j++ {
			// Calculate if number + reverse(number) is palindrom
			x = new(big.Int).Add(x, utils.Reverse(x))
			isPalindromic := utils.IsPalindromNumber(x)
			if isPalindromic {
				break
			}

			// Assume that we have found lychrel number if after maxDepth
			// iterations we stil have not found a palindrome.
			if j == maxDepth-1 {
				lychrels = append(lychrels, n)
			}
		}
	}

	// for _, v := range lychrels {
	// 	fmt.Printf("%v\n", v)
	// }
	fmt.Printf("Found %v lychrel numbers\n", len(lychrels))
}

func main() {
	N, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic("First argument missing: upper limit N")
	}
	maxDepth, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		panic("Second argument missing: maxDepth")
	}

	run(N, maxDepth)
}
