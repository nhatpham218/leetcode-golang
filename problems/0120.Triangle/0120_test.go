package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question120 struct {
	name string
	para120
	ans120
}

// para is parameter
type para120 struct {
	triangle [][]int
}

// ans is answer
type ans120 struct {
	result int
}

func Test_Problem120(t *testing.T) {
	qs := []question120{
		{
			name:    "example 1",
			para120: para120{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}},
			ans120:  ans120{11},
		},
		{
			name:    "example 2",
			para120: para120{[][]int{{-10}}},
			ans120:  ans120{-10},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 120------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans120, q.para120
			output := minimumTotal(p.triangle)
			fmt.Printf("[input]: triangle=%v       [output]:%v\n", p.triangle, output)
			assert.Equal(t, q.ans120.result, output, "minimumTotal(%v) should return %v, got %v", p.triangle, q.ans120.result, output)
		})
	}
	fmt.Printf("\n\n\n")
}
