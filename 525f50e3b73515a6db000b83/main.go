package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func CreatePhoneNumber(numbers [10]uint) string {
	ns := make([]string, 10)
	for i := range numbers {
		ns[i] = strconv.Itoa(int(numbers[i]))
	}

	return fmt.Sprintf("\"(%s) %s-%s\"", strings.Join(ns[0:3], ""), strings.Join(ns[3:6], ""), strings.Join(ns[6:], ""))
}
