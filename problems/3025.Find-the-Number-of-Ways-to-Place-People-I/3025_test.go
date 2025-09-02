package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3025 struct {
	name string
	para3025
	ans3025
}

// para is parameter
// points represents the first parameter
type para3025 struct {
	points [][]int
}

// ans is answer
// one represents the first answer
type ans3025 struct {
	one int
}

func Test_Problem3025(t *testing.T) {

	qs := []question3025{
		{
			name:     "example 1",
			para3025: para3025{[][]int{{1, 1}, {2, 2}, {3, 3}}},
			ans3025:  ans3025{0},
		},
		{
			name:     "example 2",
			para3025: para3025{[][]int{{6, 2}, {4, 4}, {2, 6}}},
			ans3025:  ans3025{2},
		},
		{
			name:     "example 3",
			para3025: para3025{[][]int{{3, 1}, {1, 3}, {1, 1}}},
			ans3025:  ans3025{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3025------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3025, q.para3025
			output := numberOfPairs(p.points)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3025.one, output, "numberOfPairs(%v) should return %d", p.points, q.ans3025.one)
		})
	}
	fmt.Printf("\n\n\n")
}
