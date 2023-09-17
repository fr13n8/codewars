package main

import "fmt"

func Decompose(n int64) []int64 {
	return decomposeHelper(n, n*n)
}

func decomposeHelper(n int64, target int64) []int64 {
	if target == 0 {
		return []int64{}
	}
	for i := n - 1; i > 0; i-- {
		diff := target - i*i
		if diff >= 0 {
			subResult := decomposeHelper(i, diff)
			if subResult != nil {
				return append(subResult, i)
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(Decompose(50))
}
