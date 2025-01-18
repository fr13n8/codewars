package main

import (
	"fmt"
	"math/big"
)

func LastDigit(as []int) int {
	var result = big.NewInt(1)
	for i := len(as) - 1; i >= 0; i-- {
		if result.Cmp(big.NewInt(4)) == -1 {
			result.Exp(big.NewInt(int64(as[i])), result, nil)
			continue
		}

		result.Exp(big.NewInt(int64(as[i])), result.Add(result.Mod(result, big.NewInt(4)), big.NewInt(4)), nil)
	}

	return int(result.Rem(result, big.NewInt(10)).Int64())
}

func main() {
	fmt.Println(LastDigit([]int{499942, 898102, 846073})) // 6
}
