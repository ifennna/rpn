package core

import (
	"fmt"
	"math"
)

type Token struct {
	Type    string
	Literal interface{}
}

const (
	NUMBER = "number"

	PLUS     = "plus"
	MINUS    = "minus"
	MULTIPLY = "multiply"
	DIVIDE   = "divide"
	NOT      = "not"
	NOTEQ    = "not equal to"
	MOD      = "modulus"
	DECR     = "decrement"
	INCR     = "increment"

	BOOLEAN = "boolean"

	RAND = "rand"

	CLRSTACK = "clear stack"
	CLRVARS  = "clear values"
	CLRALL   = "clear stack and values"

	BITAND   = "bit and"
	BITOR    = "bit or"
	BITXOR   = "bit xor"
	BITNOT   = "bit not"
	BITLEFT  = "bit shift left"
	BITRIGHT = "bit shift right"

	BOOLAND = "bool and"
	BOOLOR  = "bool or"
	BOOLXOR = "bool xor"

	LT     = "less than"
	LTOREQ = "less than or equal to"
	EQ     = "equal to"
	GT     = "greater than"
	GTOREQ = "greater than or equal to"

	ACOS = "acos"
	ASIN = "asin"
	ATAN = "atan"
	COS  = "cos"
	COSH = "cosh"
	SIN  = "sin"
	SINH = "sinh"
	TANH = "tanh"

	CEIL  = "ceiling"
	FLOOR = "floor"
	ROUND = "round"
	IP    = "integer part"
	FP    = "floating part"
	SIGN  = "push -1, 0 or 0 depending on the sign"
	ABS   = "absolute value"
	MAX   = "min"
	MIN   = "max"

	HEX = "hex mode"
	DEC = "dec mode"
	BIN = "bin mode"
	OCT = "oct mode"

	EXP  = "exponential"
	FACT = "factorial"
	SQRT = "square root"
	LN   = "natural log"
	LOG  = "logarithm"
	POW  = "raise a number to a power"

	PICK   = "pick nth item from the stack"
	REPEAT = "repeat an operation n times"
	DEPTH  = "push current stack depth"
	DROP   = "drop top item from the stack"
	DROPN  = "drop n items from the stack"
	DUP    = "duplicate top stack item"
	DUPN   = "duplicate top n stack items in order"
	ROLL   = "roll stack upwords by n"
	ROLLD  = "roll stack downwards by n"
	STACK  = "toggle stack display from horizontal to vertical"
	SWAP   = "swap top 2 stack items"

	MACRODEF = "a macro definition"
	MACRO    = "macro"
	ASSIGN   = "assign"

	HELP = "help"
	EXIT = "exit"
)

// ParseToken -> Parse a string into a calculator token
func ParseToken(item string) (Token, error) {
	var token Token
	switch item {
	case "+":
		token = makeToken(PLUS)
	case "-":
		token = makeToken(MINUS)
	case "*":
		token = makeToken(MULTIPLY)
	case "/":
		token = makeToken(DIVIDE)
	case "!":
		token = makeToken(NOT)
	case "!=":
		token = makeToken(NOTEQ)
	case "%":
		token = makeToken(MOD)
	case "--":
		token = makeToken(DECR)
	case "++":
		token = makeToken(INCR)
	case "e":
		token = Token{Type: NUMBER, Literal: math.E}
	case "pi":
		token = Token{Type: NUMBER, Literal: math.Pi}
	case "rand":
		token = makeToken(RAND)
	case "clr":
		token = makeToken(CLRSTACK)
	case "clv":
		token = makeToken(CLRVARS)
	case "cla":
		token = makeToken(CLRALL)
	case "&":
		token = makeToken(BITAND)
	case "|":
		token = makeToken(BITOR)
	case "^":
		token = makeToken(BITXOR)
	case "~":
		token = makeToken(BITNOT)
	case "<<":
		token = makeToken(BITLEFT)
	case ">>":
		token = makeToken(BITRIGHT)
	case "&&":
		token = makeToken(BOOLAND)
	case "||":
		token = makeToken(BOOLOR)
	case "^^":
		token = makeToken(BOOLXOR)
	case "<":
		token = makeToken(LT)
	case "<=":
		token = makeToken(LTOREQ)
	case "==":
		token = makeToken(EQ)
	case ">":
		token = makeToken(GT)
	case ">=":
		token = makeToken(GTOREQ)
	case "acos":
		token = makeToken(ACOS)
	case "asin":
		token = makeToken(ASIN)
	case "atan":
		token = makeToken(ATAN)
	case "cos":
		token = makeToken(COS)
	case "cosh":
		token = makeToken(COSH)
	case "sin":
		token = makeToken(SIN)
	case "sinh":
		token = makeToken(SINH)
	case "tanh":
		token = makeToken(TANH)
	case "ceil":
		token = makeToken(CEIL)
	case "floor":
		token = makeToken(FLOOR)
	case "round":
		token = makeToken(ROUND)
	case "ip":
		token = makeToken(IP)
	case "fp":
		token = makeToken(FP)
	case "sign":
		token = makeToken(SIGN)
	case "abs":
		token = makeToken(ABS)
	case "max":
		token = makeToken(MAX)
	case "min":
		token = makeToken(MIN)
	case "exp":
		token = makeToken(EXP)
	case "fact":
		token = makeToken(FACT)
	case "sqrt":
		token = makeToken(SQRT)
	case "ln":
		token = makeToken(LN)
	case "log":
		token = makeToken(LOG)
	case "pow":
		token = makeToken(POW)
	case "pick":
		token = makeToken(PICK)
	case "repeat":
		token = makeToken(REPEAT)
	case "depth":
		token = makeToken(DEPTH)
	case "drop":
		token = makeToken(DROP)
	case "dropn":
		token = makeToken(DROPN)
	case "dup":
		token = makeToken(DUP)
	case "dupn":
		token = makeToken(DUPN)
	case "roll":
		token = makeToken(ROLL)
	case "rolld":
		token = makeToken(ROLLD)
	case "stack":
		token = makeToken(STACK)
	case "swap":
		token = makeToken(SWAP)
	case "hex":
		token = makeToken(HEX)
	case "dec":
		token = makeToken(DEC)
	case "oct":
		token = makeToken(OCT)
	case "bin":
		token = makeToken(BIN)
	case "macro":
		token = makeToken(MACRODEF)
	case "x=":
		token = Token{Type: ASSIGN, Literal: "x"}
	case "exit":
		token = makeToken(EXIT)
	default:
		number, err := getInput(item)
		if err == nil {
			return Token{Type: NUMBER, Literal: number}, nil
		}
		if x, ok := values[item]; ok {
			return x, nil
		}
		if _, ok := macros[item]; ok {
			return Token{Type: MACRO, Literal: item}, nil
		}
		return Token{}, fmt.Errorf("Unknown command: %v", item)
	}
	return token, nil
}

func makeToken(tokenType string) Token {
	return Token{Type: tokenType, Literal: nil}
}
