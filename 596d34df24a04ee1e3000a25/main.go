package main

import "fmt"

func main() {
	fmt.Println(CountOnes(4, 7))
}

func CountOnes(left, right uint64) uint64 {
	return count(right+1) - count(left)
}

func countOnesOfPowersTwo(x uint64) uint64 {
	if x <= 1 {
		return 0
	}
	return 2*countOnesOfPowersTwo(x/2) + (x / 2)
}

func getNearestPowerTwo(x uint64) (near uint64) {
	if x == 0 {
		return 0
	}
	if x == 1 || x == 2 {
		return x
	}
	near = 1
	x >>= 1
	for x > 0 {
		near <<= 1
		x >>= 1
	}
	return
}

func count(x uint64) uint64 {
	if x == 0 || x == 1 {
		return 0
	}
	nearestPowerOfTwo := getNearestPowerTwo(x)
	return countOnesOfPowersTwo(nearestPowerOfTwo) + count(x-nearestPowerOfTwo) + (x - nearestPowerOfTwo)
}
