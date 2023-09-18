package main

import (
	"fmt"
	"math"
)

// https://oeis.org/A002113/b002113.txt
func main() {
	fmt.Println(FindReverseNumber(19999) == 99999999)
}

func reverseNumber(n uint64) uint64 {
	reversed := uint64(0)
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

func FindReverseNumber(n uint64) uint64 {
	l := uint64(math.Log10(float64(n)))
	var p uint64
	if 1 < n && n < uint64(1.1*math.Pow10(int(l))) {
		l = l - 1
	}
	p = uint64(math.Pow10(int(l)))

	n -= uint64(p)

	var rev uint64
	if n >= p {
		rev = n / 10
	} else {
		rev = n
	}

	return n*uint64(math.Pow10(int(l))) + reverseNumber(rev)
}
