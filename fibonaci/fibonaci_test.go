package fibonaci

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateFibonaci(t *testing.T) {
	testCases := []struct {
		startIndex int
		endIndex   int
		expected   []Fibonaci
	}{
		{
			startIndex: 0,
			endIndex:   10,
			expected: []Fibonaci{
				{Index: 0, Value: 0},
				{Index: 1, Value: 1},
				{Index: 2, Value: 1},
				{Index: 3, Value: 2},
				{Index: 4, Value: 3},
				{Index: 5, Value: 5},
				{Index: 6, Value: 8},
				{Index: 7, Value: 13},
				{Index: 8, Value: 21},
				{Index: 9, Value: 34},
				{Index: 10, Value: 55},
			},
		},
		{
			startIndex: 5,
			endIndex:   10,
			expected: []Fibonaci{
				{Index: 5, Value: 5},
				{Index: 6, Value: 8},
				{Index: 7, Value: 13},
				{Index: 8, Value: 21},
				{Index: 9, Value: 34},
				{Index: 10, Value: 55},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run("fibonaci", func(t *testing.T) {
			actualExpr := makeFibonaci(tc.startIndex, tc.endIndex)
			require.EqualValues(t, tc.expected, actualExpr)
		})
	}
}
