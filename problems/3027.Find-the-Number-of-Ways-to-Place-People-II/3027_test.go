package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3027 struct {
	name string
	para3027
	ans3027
}

// para is parameter
// points represents the first parameter
type para3027 struct {
	points [][]int
}

// ans is answer
// one represents the first answer
type ans3027 struct {
	one int
}

func Test_Problem3027(t *testing.T) {

	qs := []question3027{
		{
			name:     "example 1",
			para3027: para3027{[][]int{{1, 1}, {2, 2}, {3, 3}}},
			ans3027:  ans3027{0},
		},
		{
			name:     "example 2",
			para3027: para3027{[][]int{{6, 2}, {4, 4}, {2, 6}}},
			ans3027:  ans3027{2},
		},
		{
			name:     "example 3",
			para3027: para3027{[][]int{{3, 1}, {1, 3}, {1, 1}}},
			ans3027:  ans3027{2},
		},
		{
			name:     "two points same x",
			para3027: para3027{[][]int{{1, 5}, {1, 3}}},
			ans3027:  ans3027{1},
		},
		{
			name:     "two points same y",
			para3027: para3027{[][]int{{1, 3}, {5, 3}}},
			ans3027:  ans3027{1},
		},
		{
			name:     "negative coordinates",
			para3027: para3027{[][]int{{-1, 1}, {0, 0}, {1, -1}}},
			ans3027:  ans3027{2},
		},
		{
			name:     "large coordinates",
			para3027: para3027{[][]int{{-1000000000, 1000000000}, {0, 0}, {1000000000, -1000000000}}},
			ans3027:  ans3027{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3027------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3027, q.para3027
			output := numberOfPairs(p.points)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3027.one, output, "numberOfPairs(%v) should return %d", p.points, q.ans3027.one)
		})
	}
	fmt.Printf("\n\n\n")
}
