package Sequnces

import "testing"

func TestDuplicate(t *testing.T) {
	tests := []struct {
		input  []int
		expect bool
	}{
		{input: []int{1, 2, 3, 1}, expect: true},
		{input: []int{1, 2, 3, 4}, expect: false},
		{input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, expect: true},
	}

	for _, test := range tests {
		if out := containsDuplicate(test.input); out != test.expect {
			t.Error("expected", test.expect, "got", out, "intput", test.input)
		}
	}
}
