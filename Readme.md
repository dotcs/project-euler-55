# Project Euler #55 - Solution in Go

This code solves the question from [Project Euler #55](https://projecteuler.net/problem=55).
It's written in Golang and is a rather brute-force approach to solve the problem.

## Learning

- Numbers quickly get very large: `int64` is not sufficient, instead I had to choose [`big.Int`](https://golang.org/pkg/math/big/)
- Although this is a brute force solution it runs quite fast. On my Macbook it takes about 90ms.
- There is lot of room for improvement. Intermediate results could be cached and the task could be parallelized.