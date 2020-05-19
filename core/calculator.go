package core

import (
	"fmt"
	"os"
)

var stack []Token

// Calculate -> run a calculation for a sequence of commands
func Calculate(args []string) {
	for _, item := range args {
		token, err := ParseToken(item)
		if err != nil {
			fmt.Printf("rpn: %v\n", err)
			os.Exit(1)
		}

		switch token.Type {
		case NUMBER:
			push(token)
		case PLUS:
			op1 := pop()
			op2 := pop()
			push(Token{Type: NUMBER, Literal: op1 + op2})
		case MINUS:
			op1 := pop()
			op2 := pop()
			push(Token{Type: NUMBER, Literal: op1 - op2})
		case MULTIPLY:
			op1 := pop()
			op2 := pop()
			push(Token{Type: NUMBER, Literal: op1 * op2})
		case DIVIDE:
			op1 := pop()
			op2 := pop()
			push(Token{Type: NUMBER, Literal: op1 / op2})
		}
	}

	result := stack[0]
	fmt.Println(result.Literal)
}

func push(element Token) {
	stack = append(stack, element)
}

func pop() float64 {
	length := len(stack)
	if length == 0 {
		fmt.Println("rpn: Temporary empty stack error")
		os.Exit(1)
	}
	var element Token
	stack, element = stack[:length-1], stack[length-1]

	return element.Literal
}
