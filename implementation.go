package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

// OperandCheck checks if character is numer
func OperandCheck(character byte) bool {
	operands := "123456789"
	return strings.ContainsRune(operands, rune(character))
}

// PrefixEvaluate evlauate prefix expression
func PrefixEvaluate(input string) (string, error) {
	var stack Stack
	var res float64

	operations := map[byte]func(float64, float64) float64{
		'+': func(num1, num2 float64) float64 { return num1 + num2 },
		'-': func(num1, num2 float64) float64 { return num1 - num2 },
		'*': func(num1, num2 float64) float64 { return num1 * num2 },
		'/': func(num1, num2 float64) float64 { return num1 / num2 },
		'^': func(num1, num2 float64) float64 { return math.Pow(num1, num2) },
	}
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != ' ' {

			if OperandCheck(input[i]) {
				stack.Push(string(input[i]))
			} else {
				opFunc := operations[input[i]]
				value1, ok1 := stack.Pop()
				value2, ok2 := stack.Pop()

				if !ok1 || !ok2 {
					return "", fmt.Errorf("Pop fails")
				}

				num1, _ := strconv.ParseFloat(value1, 64)
				num2, _ := strconv.ParseFloat(value2, 64)
				res = opFunc(num1, num2)

				stack.Push(strconv.FormatFloat(res, 'f', -1, 64))
			}
		}
	}

	result, ok := stack.Pop()
	if !ok {
		return "", fmt.Errorf("Pop fails")
	}
	return result, nil
}
