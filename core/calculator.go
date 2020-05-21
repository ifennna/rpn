package core

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

var stack []Token
var values map[string]Token
var macros map[string][]Token

var mode = DEC
var display = "horizontal"

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

// Repl -> create a read-eval-print loop
func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		if scanned := scanner.Scan(); !scanned {
			return
		}
		text := strings.Split(scanner.Text(), " ")
		for _, item := range text {
			if item == "" {
				continue
			}
			token, err := ParseToken(strings.TrimSpace(item))
			if err != nil {
				fmt.Printf("rpn: %v\n", err)
				os.Exit(1)
			}

			handleCommand(token)
		}
	}
}

func printPrompt() {
	var valueStack []interface{}
	for _, item := range stack {
		valueStack = append(valueStack, item.Literal)
	}
	if display == "horizontal" {
		fmt.Printf("%v > ", valueStack)
	} else {
		fmt.Println("STACK TOP")
		for i := len(valueStack) - 1; i >= 0; i-- {
			fmt.Printf("%v\n", valueStack[i])
		}
		fmt.Println("STACK BOTTOM")
		fmt.Print("> ")
	}
}

func handleCommand(token Token) {
	switch token.Type {
	case NUMBER:
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
		op1 := popBoolean(NOT)
		push(Token{Type: BOOLEAN, Literal: !op1})
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
		push(Token{Type: NUMBER, Literal: float64(int64(op1) & int64(op2))})
	case BITOR:
		op1 := popNumber(BITOR)
		op2 := popNumber(BITOR)
		push(Token{Type: NUMBER, Literal: float64(int64(op1) | int64(op2))})
	case BITXOR:
		op1 := popNumber(BITXOR)
		op2 := popNumber(BITXOR)
		push(Token{Type: NUMBER, Literal: float64(int64(op1) ^ int64(op2))})
	case BITNOT:
		op1 := popNumber(BITNOT)
		op2 := popNumber(BITNOT)
		push(Token{Type: NUMBER, Literal: float64(int64(op1) &^ int64(op2))})
	case BITLEFT:
		op1 := popNumber(BITLEFT)
		op2 := popNumber(BITLEFT)
		push(Token{Type: NUMBER, Literal: float64(int64(op1) << int64(op2))})
	case BITRIGHT:
		op1 := popNumber(BITRIGHT)
		op2 := popNumber(BITRIGHT)
		push(Token{Type: NUMBER, Literal: float64(int64(op1) >> int64(op2))})

	case BOOLAND:
		op1 := popBoolean(BOOLAND)
		op2 := popBoolean(BOOLAND)
		push(Token{Type: BOOLEAN, Literal: op1 && op2})
	case BOOLOR:
		op1 := popBoolean(BOOLAND)
		op2 := popBoolean(BOOLAND)
		push(Token{Type: BOOLEAN, Literal: op1 || op2})
	case BOOLXOR:
		op1 := popBoolean(BOOLAND)
		op2 := popBoolean(BOOLAND)
		push(Token{Type: BOOLEAN, Literal: op1 != op2})

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
		op1 := popNumber(IP)
		push(Token{Type: NUMBER, Literal: float64(int(op1))})
	case FP:
		op1 := popNumber(IP)
		value := op1 - float64(int(op1))
		push(Token{Type: NUMBER, Literal: value})
	case SIGN:
		op1 := popNumber(IP)
		if op1 >= 0 {
			push(Token{Type: NUMBER, Literal: 0})
		} else {
			push(Token{Type: NUMBER, Literal: -1})
		}
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
		op1 := popNumber(FACT)
		push(Token{Type: NUMBER, Literal: factorial(op1)})
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
		op1 := popNumber(PICK)
		if len(stack) <= int(op1) {
			throwNotEnoughElementsError(PICK)
		}
		stack = remove(stack, int(op1))
	case REPEAT:
	case DEPTH:
		push(Token{Type: NUMBER, Literal: len(stack) - 1})
	case DROP:
		if len(stack) < 1 {
			throwNotEnoughElementsError(DUP)
		}
		pop()
	case DROPN:
		n := popNumber(DROPN)
		if len(stack) < int(n) {
			throwNotEnoughElementsError(DROPN)
		}
		for i := 0; i <= int(n); i++ {
			pop()
		}
	case DUP:
		if len(stack) < 1 {
			throwNotEnoughElementsError(DUP)
		}
		op1, _ := pop()
		push(op1)
		push(op1)
	case DUPN:
	case ROLL:
	case ROLLD:
	case STACK:
		if display == "horizontal" {
			display = "vertical"
		} else {
			display = "horizontal"
		}
	case SWAP:
		if len(stack) < 2 {
			throwNotEnoughElementsError(SWAP)
		}
		op1, _ := pop()
		op2, _ := pop()
		push(op1)
		push(op2)
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

func popBoolean(command string) bool {
	item, err := pop()
	if err != nil {
		throwNotEnoughElementsError(command)
	}
	if item.Type != BOOLEAN {
		throwWrongElementType(BOOLEAN, item.Type)
	}

	return item.Literal.(bool)
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

func factorial(n float64) float64 {
	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}

func remove(slice []Token, s int) []Token {
	return append(slice[:s], slice[s+1:]...)
}

func throwNotEnoughElementsError(action string) {
	fmt.Printf("rpn: Not enough items on the stack to perform this command: %v\n", action)
	os.Exit(1)
}

func throwWrongElementType(expected, actual string) {
	fmt.Printf("rpn: Expected a %v on the stack but found a %v\n", expected, actual)
	os.Exit(1)
}
