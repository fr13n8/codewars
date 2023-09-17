package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1234567890
	fmt.Println(NextBigger(n))
}

func NextBigger(n int) int {
	num := n
	digits := []int{}
	for num > 0 {
		digits = append([]int{num % 10}, digits...)
		num /= 10
	}
	if len(digits) == 0 || len(digits) == 1 {
		return -1
	}

	i := 0
	for i = len(digits) - 1; i > 0; i-- {
		if digits[i] > digits[i-1] {
			break
		}
	}

	if i == 0 {
		return -1
	}

	for j := len(digits) - 1; j >= i; j-- {
		if digits[i-1] < digits[j] {
			temp := digits[i-1]
			digits[i-1] = digits[j]
			digits[j] = temp
			break
		}
	}

	ls := digits[i:]
	fs := digits[:i]

	for i, j := 0, len(ls)-1; i < j; i, j = i+1, j-1 {
		ls[i], ls[j] = ls[j], ls[i]
	}

	rs := append(fs, ls...)
	rn := ""
	for i := 0; i < len(rs); i++ {
		rn += strconv.Itoa(rs[i])
	}

	result, _ := strconv.Atoi(rn)
	return result
}
