package lab2_test

import (
	"strings"

	"gopkg.in/check.v1"

	lab2 "github.com/Farid545/Team-Lab-2"
)

type ComputeHandlerSuite struct{}

var _ = check.Suite(&ComputeHandlerSuite{})

// TestAddition checks whether the result of the calculation
// corresponding to the Input is written to the Output.

func (s *ComputeHandlerSuite) TestAddition(c *check.C) {
	input := "+ 2 3"
	expectedOutput := "5"

	inputReader := strings.NewReader(input)
	var outputBuilder strings.Builder

	handler := &lab2.ComputeHandler{
		Input:  inputReader,
		Output: &outputBuilder,
	}

	err := handler.Compute()

	c.Assert(err, check.IsNil)
	c.Assert(outputBuilder.String(), check.Equals, expectedOutput)
}

// TestInvalidSyntax checks if an error is returned in case
// of a syntax problem in the Input.

func (s *ComputeHandlerSuite) TestInvalidSyntax(c *check.C) {
	input := "* 2 + 3"
	inputReader := strings.NewReader(input)
	var outputBuilder strings.Builder

	handler := &lab2.ComputeHandler{
		Input:  inputReader,
		Output: &outputBuilder,
	}

	err := handler.Compute()

	c.Assert(err, check.NotNil)
}
