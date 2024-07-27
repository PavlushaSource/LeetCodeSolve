package Sequnces

import "testing"

func Test121(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{input: []int{7, 1, 5, 3, 6, 4}, expect: 5},
		{input: []int{7, 6, 4, 3, 1}, expect: 0},
	}

	for _, test := range tests {
		if out := maxProfit(test.input); out != test.expect {
			t.Error("expected", test.expect, "got", out, "intput", test.input)
		}
	}
}
