package here

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoc(t *testing.T) {
	type testCase struct {
		input       string
		expectation string
	}

	testCases := []testCase{
		{
			input:       "Test",
			expectation: "Test",
		},
		{
			input: `
				Test
			`,
			expectation: "Test",
		},
		{
			input: `
				Test
				  - Test
			`,
			expectation: "Test\n  - Test",
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expectation, Doc(tc.input))
	}
}
