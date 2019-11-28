package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dotcs/project-euler-55/utils"
)

func run(N, maxDepth int64) {
	lychrels := make([]int64, 0)

	var n int64

	for n = 1; n <= N; n++ {
		if utils.IsLychrel(n, maxDepth) {
			lychrels = append(lychrels, n)
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
