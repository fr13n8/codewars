package main

import "fmt"

func Partitions(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j <= i {
				dp[i][j] = dp[i][j-1] + dp[i-j][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[n][n]
}

func main() {
	n := 5
	result := Partitions(n)
	fmt.Println(result)
}
