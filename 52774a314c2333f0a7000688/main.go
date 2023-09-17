package main

import "fmt"

func main() {
	fmt.Println(ValidParentheses("(())((()())())"))
}

func ValidParentheses(parens string) bool {
	// Create a stack to store the parentheses
	stack := []rune{}

	// Iterate through the string
	for _, c := range parens {
		// If the character is an opening parenthesis, push it to the stack
		if c == '(' {
			stack = append(stack, c)
		} else if c == ')' {
			// If the character is a closing parenthesis and the stack is empty, return false
			// because there is no matching opening parenthesis
			if len(stack) == 0 {
				return false
			}

			// Pop the top element from the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If the stack is not empty, return false because there are unbalanced parentheses
	if len(stack) > 0 {
		return false
	}

	// Otherwise, the parentheses are balanced
	return true
}
