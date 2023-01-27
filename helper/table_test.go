package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTable(t *testing.T) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Danu",
			request:  "Danu",
			expected: "Hello World Danu",
		},
		{
			name:     "Naruto",
			request:  "Naruto",
			expected: "Hello World Naruto",
		},
	}

	// _ ignore index
	for _, test := range tests {
		// sub test
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
