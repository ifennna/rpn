package core

import (
	"fmt"
	"strconv"
)

type Token struct {
	Type    string
	Literal float64
}

const (
	NUMBER = "number"

	PLUS     = "plus"
	MINUS    = "minus"
	MULTIPLY = "multiply"
	DIVIDE   = "divide"
)

// ParseToken -> Parse a string into a calculator token
func ParseToken(item string) (Token, error) {
	number, err := strconv.ParseFloat(item, 64)
	if err == nil {
		return Token{Type: NUMBER, Literal: number}, nil
	}

	var token Token
	switch item {
	case "+":
		token = Token{Type: PLUS, Literal: 0}
	case "-":
		token = Token{Type: MINUS, Literal: 0}
	case "*":
		token = Token{Type: MULTIPLY, Literal: 0}
	case "/":
		token = Token{Type: DIVIDE, Literal: 0}
	default:
		return Token{}, fmt.Errorf("Unknown command: %v", item)
	}
	return token, nil
}
