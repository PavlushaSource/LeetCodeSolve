package Sequnces

import (
	"slices"
	"testing"
)

func Test15(t *testing.T) {
	tests := []struct {
		input  []int
		expect [][]int
	}{
		{input: []int{3, 0, -2, -1, 1, 2}, expect: [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},
		{input: []int{-1, 0, 1, 2, -1, -4}, expect: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{input: []int{0, 1, 1}, expect: [][]int{}},
		{input: []int{0, 0, 0}, expect: [][]int{{0, 0, 0}}},
	}

	for _, test := range tests {
		out := threeSum(test.input)
		if len(out) != len(test.expect) {
			t.Fatal("expected", test.expect, "got", out)
		}
		for index := range out {
			if !slices.ContainsFunc(out, func(e []int) bool {
				return slices.Equal(e, test.expect[index])
			}) {
				t.Error("expected", test.expect, "got", out)
			}
		}
	}
}
