package studyhall

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	cases := []struct{
		name string
		input []int
		expectedOut []int
	}{
		{
			name: "normal_case",
			input: []int{9,4,7,6,5,1,2,8,3,2},
			expectedOut: []int{1,2,2,3,4,5,6,7,8,9},
		},
		{
			name: "already_sorted",
			input: []int{1,2,2,3,4,5,6,7,8,9},
			expectedOut: []int{1,2,2,3,4,5,6,7,8,9},
		},
		{
			name: "empty_slice",
			input: []int{},
			expectedOut: []int{},
		},
		{
			name: "one_element",
			input: []int{1},
			expectedOut: []int{1},
		},
		{
			name: "two_element",
			input: []int{9,1},
			expectedOut: []int{1,9},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := MergeSort(tc.input)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}
