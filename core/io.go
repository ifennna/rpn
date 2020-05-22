package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mode = DEC
var display = "horizontal"
var repl = false

// Calculate -> run a calculation for a sequence of commands
func Calculate(args []string) {
	var input []string
	// check if there's anything in stdin (from a pipe perhaps)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		if scanned := scanner.Scan(); scanned {
			input = strings.Split(scanner.Text(), " ")
		}
	}

	eval(append(args, input...))

	if len(stack) > 0 {
		result := stack[0]
		fmt.Fprintln(os.Stdout, result.Literal)
	}
}

// Repl -> create a read-eval-print loop
func Repl() {
	repl = true
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		if scanned := scanner.Scan(); !scanned {
			return
		}
		text := strings.Split(scanner.Text(), " ")
		eval(text)
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

func throwNotEnoughElementsError(action string) {
	fmt.Fprintf(os.Stderr, "rpn: Not enough items on the stack to perform this command: %v\n", action)
	exit()
}

func throwNotEnoughArgumentsError(action string) {
	fmt.Fprintf(os.Stderr, "rpn: Not enough arguments to perform this command: %v\n", action)
	exit()
}

func throwWrongElementType(expected, actual string) {
	fmt.Fprintf(os.Stderr, "rpn: Expected a %v on the stack but found a %v\n", expected, actual)
	exit()
}

func exit() {
	if !repl {
		os.Exit(1)
	}
}
