package numbers_test

import (
	"fmt"
	"slices"
	"testing"

	numbers "github.com/wizardx-develop/learn-github-actions-on-example/numbers"
)

func TestFindPrimeNumbers(t *testing.T) {
	var tests = []struct {
		in   float64
		want []float64
	}{
		{0, []float64{}},
		{1, []float64{}},
		{2, []float64{2}},
		{10, []float64{2, 3, 5, 7}},
		{100, []float64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%f", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := numbers.FindPrimeNumbers(tt.in)
			// if err != nil {
			// 	t.Errorf("Got error: %v", err)
			// }
			if slices.Compare(tt.want, ans) != 0 {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}
}
