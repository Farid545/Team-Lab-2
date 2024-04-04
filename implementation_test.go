package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type PrefixEvaluateSuite struct{}

var _ = Suite(&PrefixEvaluateSuite{})

func (s *PrefixEvaluateSuite) TestAdditionOfTwoOperands(c *C) {
	input := "+ 2 3"
	expected := "5"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestSubtractionOfTwoOperands(c *C) {
	input := "- 5 3"
	expected := "2"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestMultiplicationOfTwoOperands(c *C) {
	input := "* 2 3"
	expected := "6"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestDivisionOfTwoOperands(c *C) {
	input := "/ 6 2"
	expected := "3"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestExponentiationOfTwoOperands(c *C) {
	input := "^ 2 3"
	expected := "8"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestAdditionOfThreeOperands(c *C) {
	input := "+ + 1 2 3"
	expected := "6"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestMultiplicationOfThreeOperands(c *C) {
	input := "* + 1 2 3"
	expected := "9"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestSubtractionOfThreeOperands(c *C) {
	input := "- * 3 4 2"
	expected := "10"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestExponentiationOfThreeOperands(c *C) {
	input := "^ / 9 3 2"
	expected := "9"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestCombinationOfOperatorsAndOperands(c *C) {
	input := "+ * 2 3 5"
	expected := "11"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestEmptyInputString(c *C) {
	input := ""
	_, err := PrefixEvaluate(input)
	c.Assert(err, NotNil)
}

func (s *PrefixEvaluateSuite) TestInvalidCharacters(c *C) {
	input := "+ a b"
	_, err := PrefixEvaluate(input)
	c.Assert(err, NotNil)
}
func (s *PrefixEvaluateSuite) TestEightOperands1(c *C) {
	input := "+/-^23*/4232/4^22"
	expected := "2"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

func (s *PrefixEvaluateSuite) TestSevenOperands1(c *C) {
	input := "+-*+12^22 ^41 /2 1"
	expected := "10"
	result, err := PrefixEvaluate(input)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, expected)
}

// Example of using PrefixEvaluate function
func ExamplePrefixEvaluate() {
	expressions := []string{
		"+ 2 3",
		"- 5 3",
		"* 2 3",
		"/ 6 2",
		"^ 2 3",
		"+ + 1 2 3",
		"* + 1 2 3",
	}

	for _, expr := range expressions {
		result, err := PrefixEvaluate(expr)
		if err != nil {
			fmt.Printf("Error evaluating expression '%s': %s\n", expr, err)
		} else {
			fmt.Printf("Result of expression '%s' is %s\n", expr, result)
		}
	}

	// Output:
	// Result of expression '+ 2 3' is 5
	// Result of expression '- 5 3' is 2
	// Result of expression '* 2 3' is 6
	// Result of expression '/ 6 2' is 3
	// Result of expression '^ 2 3' is 8
	// Result of expression '+ + 1 2 3' is 6
	// Result of expression '* + 1 2 3' is 9
}
