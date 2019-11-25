package main

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/dotcs/project-euler-55/utils"
)

func run(N, maxDepth int) {
	lychrels := make([]int64, 0)

	for n := 1; n <= N; n++ {
		x := new(big.Int)
		x.SetString(strconv.FormatInt(int64(n), 10), 10)

		for j := 0; j < maxDepth; j++ {
			// Calculate if number + reverse(number) is palindrom
			x = new(big.Int).Add(x, utils.Reverse(x))
			isPalindromic := utils.IsPalindromNumber(*x)
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
	N := 10000
	maxDepth := 50

	run(N, maxDepth)
}
