package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3660 struct {
	name string
	para3660
	ans3660
}

// para is parameter
// one represents the first parameter
type para3660 struct {
	one []int
}

// ans is answer
// one represents the first answer
type ans3660 struct {
	one []int
}

func Test_Problem3660(t *testing.T) {

	qs := []question3660{
		{
			name:     "example 1",
			para3660: para3660{[]int{2, 1, 3}},
			ans3660:  ans3660{[]int{2, 2, 3}},
		},
		{
			name:     "example 2",
			para3660: para3660{[]int{2, 3, 1}},
			ans3660:  ans3660{[]int{3, 3, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3660------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3660, q.para3660
			output := maxValue(p.one)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3660.one, output, "maxValue(%v) should return %v, got %v", p.one, q.ans3660.one, output)
		})
	}
	fmt.Printf("\n\n\n")
}
