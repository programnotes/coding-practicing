package algorithms

import (
	"testing"
)

var _searchTestCases = []SearchTestCase{
	{
		values:   []int{},
		target:   100,
		expected: -1,
	},
	{
		values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		target:   1,
		expected: 0,
	},
	{
		values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		target:   100,
		expected: -1,
	},
	{
		values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		target:   5,
		expected: 4,
	},
}

func TestBinary(t *testing.T) {
	for _, val := range _searchTestCases {
		index := BinarySearch(val.values, val.target)
		if index != val.expected {
			t.Errorf("BinarySearch(%v, %v) = %v, expected %v",
				val.values, val.target, index, val.expected)
		}
	}
}