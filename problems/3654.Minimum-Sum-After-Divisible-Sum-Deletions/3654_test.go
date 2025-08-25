package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3654 struct {
	name string
	para3654
	ans3654
}

// para is parameter
// nums represents the first parameter
// k represents the second parameter
type para3654 struct {
	nums []int
	k    int
}

// ans is answer
// one represents the first answer
type ans3654 struct {
	one int64
}

func Test_Problem3654(t *testing.T) {

	qs := []question3654{
		{
			name:     "example 1",
			para3654: para3654{[]int{1, 1, 1}, 2},
			ans3654:  ans3654{1},
		},
		{
			name:     "example 2",
			para3654: para3654{[]int{3, 1, 4, 1, 5}, 3},
			ans3654:  ans3654{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3654------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3654, q.para3654
			output := minArraySum(p.nums, p.k)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3654.one, output, "minArraySum(%v, %d) should return %d", p.nums, p.k, q.ans3654.one)
		})
	}
	fmt.Printf("\n\n\n")
}
