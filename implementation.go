package lab2

import (
	"strconv"
	"strings"
)

type Stack []string
type operationFunc func(int, int) int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func OperandCheck(character byte) bool {
	operands := "123456789"
	return strings.ContainsRune(operands, rune(character))
}

// TODO: document this function.
// PrefixEvaluate converts
func PrefixEvaluate(input string) (string, error) {
	var stack Stack
	var res int

	operations := map[byte]operationFunc{
		'+': func(num1, num2 int) int { return num1 + num2 },
		'-': func(num1, num2 int) int { return num1 - num2 },
		'*': func(num1, num2 int) int { return num1 * num2 },
		'/': func(num1, num2 int) int { return num1 / num2 },
	}

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != ' ' {
			if OperandCheck(input[i]) {
				stack.Push(string(input[i]))
			} else {
				value1, _ := stack.Pop()
				value2, _ := stack.Pop()

				num1, _ := strconv.Atoi(value1)
				num2, _ := strconv.Atoi(value2)

				if opFunc, exists := operations[input[i]]; exists {
					res = opFunc(num1, num2)
					stack.Push(strconv.Itoa(res))
				}

			}
		}

	}

	return strconv.Itoa(res), nil
}
