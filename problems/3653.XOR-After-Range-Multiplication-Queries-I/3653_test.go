package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3653 struct {
	name string
	para3653
	ans3653
}

// para is parameter
// nums represents the first parameter
// queries represents the second parameter
type para3653 struct {
	nums    []int
	queries [][]int
}

// ans is answer
// one represents the first answer
type ans3653 struct {
	one int
}

func Test_Problem3653(t *testing.T) {

	qs := []question3653{
		{
			name:     "example 1",
			para3653: para3653{[]int{1, 1, 1}, [][]int{{0, 2, 1, 4}}},
			ans3653:  ans3653{4},
		},
		{
			name:     "example 2",
			para3653: para3653{[]int{2, 3, 1, 5, 4}, [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}}},
			ans3653:  ans3653{31},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3653------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3653, q.para3653
			output := xorAfterQueries(p.nums, p.queries)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3653.one, output, "xorAfterQueries(%v, %v) should return %d", p.nums, p.queries, q.ans3653.one)
		})
	}
	fmt.Printf("\n\n\n")
}
