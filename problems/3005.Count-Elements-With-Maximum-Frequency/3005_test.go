package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3005 struct {
	name string
	para3005
	ans3005
}

// para is parameter
// one represents the first parameter
type para3005 struct {
	one []int
}

// ans is answer
// one represents the first answer
type ans3005 struct {
	one int
}

func Test_Problem3005(t *testing.T) {

	qs := []question3005{
		{
			name:     "example 1",
			para3005: para3005{[]int{1, 2, 2, 3, 1, 4}},
			ans3005:  ans3005{4},
		},
		{
			name:     "example 2",
			para3005: para3005{[]int{1, 2, 3, 4, 5}},
			ans3005:  ans3005{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3005------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3005, q.para3005
			output := maxFrequencyElements(p.one)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3005.one, output, "maxFrequencyElements(%v) should return %v, got %v", p.one, q.ans3005.one, output)
		})
	}
	fmt.Printf("\n\n\n")
}
