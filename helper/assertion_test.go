package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// assert, throw t.fail
// require, throw t.failNow

func TestAssertion(t *testing.T) {
	result := HelloWorld("Danu")
	// if fail, throw t.Fail
	assert.Equal(t, "Hi Danu", result, "Result must be Hello World Danu")
	fmt.Println("Test using assertion assert")
}

func TestAssertion2(t *testing.T) {
	result := HelloWorld("Danu")
	// if fail, throw t.FailNow
	require.Equal(t, "Hi Danu", result, "Result must be Hello World Danu")
	fmt.Println("Test using assertion require")
}
