package lab2

import (
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	expression, err := ch.readExpression()
	if err != nil {
		return err
	}

	result, err := PrefixEvaluate(expression)
	if err != nil {
		return err
	}

	_, err = io.WriteString(ch.Output, result)
	if err != nil {
		return err
	}

	return nil
}

func (ch *ComputeHandler) readExpression() (string, error) {
	buffer := new(strings.Builder)
	_, err := io.Copy(buffer, ch.Input)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(buffer.String()), nil
}
