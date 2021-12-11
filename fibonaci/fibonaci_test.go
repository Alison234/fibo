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
				{index: 0, value: 0},
				{index: 1, value: 1},
				{index: 2, value: 1},
				{index: 3, value: 2},
				{index: 4, value: 3},
				{index: 5, value: 5},
				{index: 6, value: 8},
				{index: 7, value: 13},
				{index: 8, value: 21},
				{index: 9, value: 34},
				{index: 10, value: 55},
			},
		},
		{
			startIndex: 5,
			endIndex:   10,
			expected: []Fibonaci{
				{index: 5, value: 5},
				{index: 6, value: 8},
				{index: 7, value: 13},
				{index: 8, value: 21},
				{index: 9, value: 34},
				{index: 10, value: 55},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run("fibonaci", func(t *testing.T) {
			actualExpr, err := MakeFibonaci(tc.startIndex, tc.endIndex)
			require.NoError(t, err)
			require.EqualValues(t, tc.expected, actualExpr)
		})
	}

}
