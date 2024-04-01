package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixEvaluate(t *testing.T) {
	res, err := PrefixEvaluate("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "11", res)
	}
}

func ExamplePrefixEvaluate() {
	res, _ := PrefixEvaluate("+ 2 2")
	fmt.Println(res)

	// Output:
	// 4
}
