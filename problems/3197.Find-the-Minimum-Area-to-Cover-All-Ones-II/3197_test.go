package leetcode

import (
	"fmt"
	"testing"
)

type question3197 struct {
	name string
	para3197
	ans3197
}

// para is parameter
// one represents the first parameter
type para3197 struct {
	grid [][]int
}

// ans is answer
// one represents the first answer
type ans3197 struct {
	one int
}

func Test_Problem3197(t *testing.T) {

	qs := []question3197{
		{
			name:     "example 1",
			para3197: para3197{[][]int{{1, 0, 1}, {1, 1, 1}}},
			ans3197:  ans3197{5},
		},
		{
			name:     "example 2",
			para3197: para3197{[][]int{{1, 0, 1, 0}, {0, 1, 0, 1}}},
			ans3197:  ans3197{5},
		},
		{
			name:     "example 3",
			para3197: para3197{[][]int{{0, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 1, 0, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}}},
			ans3197:  ans3197{6},
		},
		{
			name:     "example 4",
			para3197: para3197{[][]int{{0, 1}, {1, 1}}},
			ans3197:  ans3197{3},
		},
		{
			name:     "example 5",
			para3197: para3197{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 1}, {1, 1, 0}}},
			ans3197:  ans3197{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3197------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3197, q.para3197
			output := minimumSum(p.grid)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			if output != q.ans3197.one {
				t.Errorf("Expected %v, got %v", q.ans3197.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
