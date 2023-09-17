package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SquareSumsRow(15))
}

var (
	found  bool
	result []int
	used   []bool
)

func isSquare(n int) bool {
	return n > 0 && math.Mod((math.Sqrt(float64(n))/1), 1.0) == 0
}

func dfs(n, step int) {
	if step == n+1 {
		found = true
		return
	}
	for i := 1; i <= n; i++ {
		if step == 1 || used[i] == false && isSquare(result[len(result)-1]+i) {
			result, used[i] = append(result, i), true
			dfs(n, step+1)
			if found {
				return
			}
			result, used[i] = result[:len(result)-1], false
		}
	}
}

func SquareSumsRow(n int) []int {
	found, result, used = false, make([]int, 0), make([]bool, n+1)
	dfs(n, 1)
	if len(result) != n {
		return nil
	}
	return result
}
