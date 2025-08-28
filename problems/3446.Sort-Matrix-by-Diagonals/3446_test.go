package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3446 struct {
	name string
	para3446
	ans3446
}

// para is parameter
type para3446 struct {
	grid [][]int
}

// ans is answer
type ans3446 struct {
	result [][]int
}

func Test_Problem3446(t *testing.T) {

	qs := []question3446{
		{
			name: "example 1",
			para3446: para3446{
				grid: [][]int{
					{1, 7, 3},
					{9, 8, 2},
					{4, 5, 6},
				},
			},
			ans3446: ans3446{
				result: [][]int{
					{8, 2, 3},
					{9, 6, 7},
					{4, 5, 1},
				},
			},
		},
		{
			name: "example 2",
			para3446: para3446{
				grid: [][]int{
					{0, 1},
					{1, 2},
				},
			},
			ans3446: ans3446{
				result: [][]int{
					{2, 1},
					{1, 0},
				},
			},
		},
		{
			name: "example 3",
			para3446: para3446{
				grid: [][]int{
					{1},
				},
			},
			ans3446: ans3446{
				result: [][]int{
					{1},
				},
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3446------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3446, q.para3446
			output := sortMatrix(p.grid)
			fmt.Printf("[input]: grid=%v       [output]:%v\n", p.grid, output)
			assert.Equal(t, q.ans3446.result, output, "sortMatrix(%v) should return %v", p.grid, q.ans3446.result)
		})
	}
	fmt.Printf("\n\n\n")
}
