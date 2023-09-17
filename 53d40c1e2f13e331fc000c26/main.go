package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(fib(2000000))
}

func multiply(a, b [][]*big.Int) [][]*big.Int {
	c := make([][]*big.Int, 2)
	for i := 0; i < 2; i++ {
		c[i] = make([]*big.Int, 2)
		for j := 0; j < 2; j++ {
			c[i][j] = new(big.Int)
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				temp := new(big.Int)
				temp.Mul(a[i][k], b[k][j])
				c[i][j].Add(c[i][j], temp)
			}
		}
	}

	return c
}

func matrixPower(a [][]*big.Int, n int64) [][]*big.Int {
	if n == 1 {
		return a
	}

	if n%2 == 0 {
		halfPower := matrixPower(a, n/2)
		return multiply(halfPower, halfPower)
	}

	temp := matrixPower(a, (n-1)/2)
	halfPower := multiply(temp, temp)
	return multiply(halfPower, a)
}

func fib(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	} else if n == 1 || n == -1 {
		return big.NewInt(1)
	}

	var negative bool
	if n < 0 {
		negative = true
		n = -n
	}

	baseMatrix := [][]*big.Int{{big.NewInt(1), big.NewInt(1)}, {big.NewInt(1), big.NewInt(0)}}
	resultMatrix := matrixPower(baseMatrix, n-1)

	if negative && n%2 == 0 {
		return resultMatrix[0][0].Neg(resultMatrix[0][0])
	}

	return resultMatrix[0][0]
}
