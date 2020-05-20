package core

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

var stack []Token
var values map[string]Token
var macros map[string][]Token

var mode = DEC

// Calculate -> run a calculation for a sequence of commands
func Calculate(args []string) {
	for _, item := range args {
		token, err := ParseToken(item)
		if err != nil {
			fmt.Printf("rpn: %v\n", err)
			os.Exit(1)
		}

		handleCommand(token)
	}

	result := stack[0]
	fmt.Println(result.Literal)
}

func handleCommand(token Token) {
	switch token.Type {
	case NUMBER, PI, E:
		push(token)
	case RAND:
		rand.Seed(time.Now().UnixNano())
		push(Token{Type: NUMBER, Literal: rand.Float64()})
	case PLUS:
		op1 := popNumber(PLUS)
		op2 := popNumber(PLUS)
		push(Token{Type: NUMBER, Literal: op1 + op2})
	case MINUS:
		op1 := popNumber(MINUS)
		op2 := popNumber(MINUS)
		push(Token{Type: NUMBER, Literal: op1 - op2})
	case MULTIPLY:
		op1 := popNumber(MULTIPLY)
		op2 := popNumber(MULTIPLY)
		push(Token{Type: NUMBER, Literal: op1 * op2})
	case DIVIDE:
		op1 := popNumber(DIVIDE)
		op2 := popNumber(DIVIDE)
		push(Token{Type: NUMBER, Literal: op1 / op2})
	case CLRSTACK:
		stack = make([]Token, 0)
	case NOT:
	case MOD:
		op1 := popNumber(MOD)
		op2 := popNumber(MOD)
		push(Token{Type: NUMBER, Literal: math.Mod(op1, op2)})
	case DECR:
		op1 := popNumber(DECR)
		push(Token{Type: NUMBER, Literal: op1 - 1})
	case INCR:
		op1 := popNumber(INCR)
		push(Token{Type: NUMBER, Literal: op1 + 1})

	case BITAND:
		op1 := popNumber(BITAND)
		op2 := popNumber(BITAND)
		push(Token{Type: NUMBER, Literal: int64(op1) & int64(op2)})
	case BITOR:
		op1 := popNumber(BITOR)
		op2 := popNumber(BITOR)
		push(Token{Type: NUMBER, Literal: int64(op1) | int64(op2)})
	case BITXOR:
		op1 := popNumber(BITXOR)
		op2 := popNumber(BITXOR)
		push(Token{Type: NUMBER, Literal: int64(op1) ^ int64(op2)})
	case BITNOT:
		op1 := popNumber(BITNOT)
		op2 := popNumber(BITNOT)
		push(Token{Type: NUMBER, Literal: int64(op1) &^ int64(op2)})
	case BITLEFT:
		op1 := popNumber(BITLEFT)
		op2 := popNumber(BITLEFT)
		push(Token{Type: NUMBER, Literal: int64(op1) << int64(op2)})
	case BITRIGHT:
		op1 := popNumber(BITRIGHT)
		op2 := popNumber(BITRIGHT)
		push(Token{Type: NUMBER, Literal: int64(op1) >> int64(op2)})

	case BOOLAND:
	case BOOLOR:
	case BOOLXOR:

	case LT:
		op1 := popNumber(LT)
		op2 := popNumber(LT)
		push(Token{Type: BOOLEAN, Literal: op1 < op2})
	case LTOREQ:
		op1 := popNumber(LT)
		op2 := popNumber(LT)
		push(Token{Type: BOOLEAN, Literal: op1 <= op2})
	case NOTEQ:
		op1 := popNumber(LT)
		op2 := popNumber(LT)
		push(Token{Type: BOOLEAN, Literal: op1 != op2})
	case EQ:
		op1 := popNumber(LT)
		op2 := popNumber(LT)
		push(Token{Type: BOOLEAN, Literal: op1 == op2})
	case GT:
		op1 := popNumber(LT)
		op2 := popNumber(LT)
		push(Token{Type: BOOLEAN, Literal: op1 > op2})
	case GTOREQ:
		op1 := popNumber(LT)
		op2 := popNumber(LT)
		push(Token{Type: BOOLEAN, Literal: op1 >= op2})

	case ACOS:
		op1 := popNumber(ACOS)
		push(Token{Type: NUMBER, Literal: math.Acos(op1)})
	case ASIN:
		op1 := popNumber(ASIN)
		push(Token{Type: NUMBER, Literal: math.Asin(op1)})
	case ATAN:
		op1 := popNumber(ATAN)
		push(Token{Type: NUMBER, Literal: math.Atan(op1)})
	case COS:
		op1 := popNumber(COS)
		push(Token{Type: NUMBER, Literal: math.Cos(op1)})
	case COSH:
		op1 := popNumber(COSH)
		push(Token{Type: NUMBER, Literal: math.Cosh(op1)})
	case SIN:
		op1 := popNumber(SIN)
		push(Token{Type: NUMBER, Literal: math.Sin(op1)})
	case SINH:
		op1 := popNumber(SINH)
		push(Token{Type: NUMBER, Literal: math.Sinh(op1)})
	case TANH:
		op1 := popNumber(TANH)
		push(Token{Type: NUMBER, Literal: math.Tanh(op1)})

	case CEIL:
		op1 := popNumber(CEIL)
		push(Token{Type: NUMBER, Literal: math.Ceil(op1)})
	case FLOOR:
		op1 := popNumber(FLOOR)
		push(Token{Type: NUMBER, Literal: math.Floor(op1)})
	case ROUND:
		op1 := popNumber(ROUND)
		push(Token{Type: NUMBER, Literal: math.Round(op1)})
	case IP:
	case FP:
	case SIGN:
	case ABS:
		op1 := popNumber(ABS)
		push(Token{Type: NUMBER, Literal: math.Abs(op1)})
	case MAX:
		op1 := popNumber(POW)
		op2 := popNumber(POW)
		push(Token{Type: NUMBER, Literal: math.Max(op1, op2)})
	case MIN:
		op1 := popNumber(POW)
		op2 := popNumber(POW)
		push(Token{Type: NUMBER, Literal: math.Min(op1, op2)})

	case HEX:
		mode = HEX
	case DEC:
		mode = DEC
	case BIN:
		mode = BIN
	case OCT:
		mode = OCT

	case EXP:
		op1 := popNumber(EXP)
		push(Token{Type: NUMBER, Literal: math.Exp(op1)})
	case FACT:
		// op1 := popNumber(FACT)
		// push(Token{Type: NUMBER, Literal: math.(op1)})
	case SQRT:
		op1 := popNumber(SQRT)
		push(Token{Type: NUMBER, Literal: math.Sqrt(op1)})
	case LN:
		op1 := popNumber(LN)
		push(Token{Type: NUMBER, Literal: math.Log(op1)})
	case LOG:
		op1 := popNumber(LOG)
		push(Token{Type: NUMBER, Literal: math.Log10(op1)})
	case POW:
		op1 := popNumber(POW)
		op2 := popNumber(POW)
		push(Token{Type: NUMBER, Literal: math.Pow(op1, op2)})

	case HNL:
	case HNS:
	case NHL:
	case NHS:

	case PICK:
	case REPEAT:
	case DEPTH:
	case DROP:
	case DROPN:
	case DUP:
	case DUPN:
	case ROLL:
	case ROLLD:
	case STACK:
	case SWAP:

	case MACRO:
	case ASSIGN:

	case HELP:
	case EXIT:
		fmt.Println("Goodbye")
		os.Exit(0)
	}
}

func push(element Token) {
	stack = append(stack, element)
}

func popNumber(command string) float64 {
	item, err := pop()
	if err != nil {
		throwNotEnoughElementsError(command)
	}
	if item.Type != NUMBER {
		throwWrongElementType(NUMBER, item.Type)
	}

	return item.Literal.(float64)
}

func pop() (Token, error) {
	length := len(stack)
	if length == 0 {
		return Token{}, fmt.Errorf("Popping from an empty stack")
	}
	var element Token
	stack, element = stack[:length-1], stack[length-1]

	return element, nil
}

func throwNotEnoughElementsError(action string) {
	fmt.Printf("rpn: Not enough items on the stack to perform this command: %v\n", action)
	os.Exit(1)
}

func throwWrongElementType(expected, actual string) {
	fmt.Printf("rpn: Expected a %v on the stack but found a %v\n", expected, actual)
	os.Exit(1)
}
