package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question611 struct {
	name string
	para611
	ans611
}

// para is parameter
type para611 struct {
	nums []int
}

// ans is answer
type ans611 struct {
	result int
}

func Test_Problem611(t *testing.T) {
	qs := []question611{
		{
			name:    "example 1",
			para611: para611{[]int{2, 2, 3, 4}},
			ans611:  ans611{3},
		},
		{
			name:    "example 2",
			para611: para611{[]int{4, 2, 3, 4}},
			ans611:  ans611{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 611------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans611, q.para611
			output := triangleNumber(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			assert.Equal(t, q.ans611.result, output, "triangleNumber(%v) should return %v, got %v", p.nums, q.ans611.result, output)
		})
	}
	fmt.Printf("\n\n\n")
}
