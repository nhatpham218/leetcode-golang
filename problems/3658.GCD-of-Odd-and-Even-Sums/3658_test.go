package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3658 struct {
	name string
	para3658
	ans3658
}

// para is parameter
// one represents the first parameter
type para3658 struct {
	one int
}

// ans is answer
// one represents the first answer
type ans3658 struct {
	one int
}

func Test_Problem3658(t *testing.T) {

	qs := []question3658{
		{
			name:     "example 1",
			para3658: para3658{4},
			ans3658:  ans3658{4},
		},
		{
			name:     "example 2",
			para3658: para3658{5},
			ans3658:  ans3658{5},
		},
		{
			name:     "n = 1",
			para3658: para3658{1},
			ans3658:  ans3658{1},
		},
		{
			name:     "n = 2",
			para3658: para3658{2},
			ans3658:  ans3658{2},
		},
		{
			name:     "n = 3",
			para3658: para3658{3},
			ans3658:  ans3658{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3658------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3658, q.para3658
			output := gcdOfOddEvenSums(p.one)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3658.one, output, "gcdOfOddEvenSums(%d) should return %d", p.one, q.ans3658.one)
		})
	}
	fmt.Printf("\n\n\n")
}
