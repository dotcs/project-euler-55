package main

import (
	"fmt"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/dotcs/project-euler-55/utils"
)

type lychrelList struct {
	entries []int64
	mux     sync.Mutex
}

func (ll *lychrelList) Append(n int64) {
	ll.mux.Lock()
	ll.entries = append(ll.entries, n)
	ll.mux.Unlock()
}

func work(ll *lychrelList, wg *sync.WaitGroup, start, stop int64, maxDepth int64) {
	defer wg.Done()

	var n, j int64
	for n = start; n <= stop; n++ {
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
				ll.Append(n)
			}
		}
	}
}

func run(N, maxDepth, cores int64) {
	lychrels := lychrelList{entries: make([]int64, 0)}

	wg := &sync.WaitGroup{}
	chunkSize := N / cores

	fmt.Printf("Start calculation on %v goroutines.\n", cores)
	var i, start, stop int64
	for i = 0; i < cores; i++ {
		wg.Add(1)
		start = i*chunkSize + 1
		stop = (i + 1) * chunkSize
		fmt.Printf("Start goroutine. Calculate numbers: %v to %v\n", start, stop)
		go work(&lychrels, wg, start, stop, maxDepth)
	}
	wg.Wait()

	// for _, v := range lychrels {
	// 	fmt.Printf("%v\n", v)
	// }
	fmt.Printf("Found %v lychrel numbers\n", len(lychrels.entries))
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
	var cores int64 = int64(runtime.NumCPU())

	run(N, maxDepth, cores)
}
