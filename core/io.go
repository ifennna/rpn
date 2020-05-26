package core

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-homedir"
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

	if config, err := getConfig(); err == nil {
		args = append(strings.Split(config, " "), args...)
	}

	eval(append(args, input...))

	if len(stack) > 0 {
		result := stack[0]
		fmt.Fprintln(os.Stdout, showResultValue(result.Literal))
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
		if config, err := getConfig(); err == nil {
			text = append(strings.Split(config, " "), text...)
		}
		eval(text)
	}
}

func getConfig() (string, error) {
	home, err := homedir.Dir()
	data, err := ioutil.ReadFile(fmt.Sprintf("%v/.rpnrc", home))
	return string(data), err
}

func printPrompt() {
	var valueStack []interface{}
	for _, item := range stack {
		valueStack = append(valueStack, item.Literal)
	}
	if display == "horizontal" {
		printRegisterValues()
		for _, item := range valueStack {
			fmt.Printf("%v ", showResultValue(item))
		}
		fmt.Print("> ")
	} else {
		fmt.Println("STACK TOP")
		for i := len(valueStack) - 1; i >= 0; i-- {
			fmt.Printf("%v\n", showResultValue(valueStack[i]))
		}
		fmt.Println("STACK BOTTOM")
		printRegisterValues()
		fmt.Print("> ")
	}
}

func printRegisterValues() {
	if len(values) > 0 {
		fmt.Print("[")
	}
	for i, v := range values {
		fmt.Printf("%v= %v", i, showResultValue(v.Literal))
	}
	if len(values) > 0 {
		fmt.Print("] ")
	}
}

func showResultValue(result interface{}) interface{} {
	if _, ok := result.(bool); ok {
		return result
	}
	switch mode {
	case DEC:
		return strconv.FormatFloat(result.(float64), 'f', -1, 64)
	case BIN:
		return getBinary(result.(float64))
	case OCT:
		return getOctal(result.(float64))
	case HEX:
		return getHex(result.(float64))
	}
	return result
}

func getInput(item string) (float64, error) {
	switch mode {
	case BIN:
		result, err := strconv.ParseInt(item, 2, 64)
		if err != nil {
			return 0, err
		}
		return float64(result), err
	case OCT:
		result, err := strconv.ParseInt(item, 8, 64)
		if err != nil {
			return 0, err
		}
		return float64(result), err
	case HEX:
		result, err := strconv.ParseInt(item, 16, 64)
		if err != nil {
			return 0, err
		}
		return float64(result), err
	default:
		return strconv.ParseFloat(item, 64)
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
	os.Exit(1)
}

func getBinary(num float64) string {
	integer := int(num)
	intPart := fmt.Sprintf("%b", integer)

	float := num - float64(integer)
	floatPart := getFloatingPart(float, 2)

	if num, _ := strconv.ParseInt(floatPart, 2, 64); num == 0 {
		return intPart
	}

	return fmt.Sprintf("%v.%v", intPart, floatPart)
}

func getOctal(num float64) string {
	integer := int(num)
	intPart := fmt.Sprintf("%o", integer)

	float := num - float64(integer)
	floatPart := getFloatingPart(float, 8)

	if num, _ := strconv.ParseInt(floatPart, 8, 64); num == 0 {
		return intPart
	}

	return fmt.Sprintf("%v.%v", intPart, floatPart)
}

func getHex(num float64) string {
	integer := int(num)
	intPart := fmt.Sprintf("%x", integer)

	float := num - float64(integer)
	floatPart := getFloatingPart(float, 16)

	if num, _ := strconv.ParseInt(floatPart, 16, 64); num == 0 {
		return intPart
	}

	return fmt.Sprintf("%v.%v", intPart, floatPart)
}

func getFloatingPart(float float64, base int) string {
	if float < 0 {
		float = math.Abs(float)
	}
	floatPart := ""
	for float != 0 {
		float *= float64(base)
		fractBit := int64(float)

		if fractBit > 0 {
			float -= float64(fractBit)
			floatPart += strconv.FormatInt(fractBit, base)
		} else {
			floatPart += "0"
		}
	}

	return floatPart
}
