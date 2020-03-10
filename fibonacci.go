package main

import (
	"fmt"
	"time"
)

func main() {
	limit := int64(45)

	var sum int64
	start := time.Now()
	for i := int64(1); i <= limit; i++ {
		sum += fibNaive(i)
	}
	end := time.Now()
	naiveTime := end.Sub(start)

	fmt.Println("Result was ", sum)
	fmt.Println("Time: ", naiveTime)

	sum = 0
	start = time.Now()
	for i := int64(1); i <= limit; i++ {
		sum += fibBetter(i, make(map[int64]int64))
	}
	end = time.Now()
	betterTime := end.Sub(start)

	fmt.Println("Result was ", sum)
	fmt.Println("Time: ", betterTime)

	fmt.Println("Better was ", naiveTime - betterTime, "better than naive")
}

/**
 A very naive fibonacci sequence generator.
 */
func fibNaive(n int64) int64 {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fibNaive(n-1) + fibNaive(n-2)
	}
}

/**
 Uses memoization and recursion to generate fibonacci sequence.
 Sacrifices memory for runtime.
 */
func fibBetter(n int64, memo map[int64]int64) int64 {
	memo[1] = 1
	memo[2] = 1

	maybeValue, isPresent := memo[n]
	if isPresent {
		return maybeValue
	} else {
		value := fibBetter(n-1, memo) + fibBetter(n-2, memo)
		memo[n] = value
		return value
	}
}
