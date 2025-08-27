package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3459 struct {
	name string
	para3459
	ans3459
}

// para is parameter
type para3459 struct {
	grid [][]int
}

// ans is answer
type ans3459 struct {
	result int
}

func Test_Problem3459(t *testing.T) {

	qs := []question3459{
		{
			name: "example 1",
			para3459: para3459{
				grid: [][]int{
					{2, 2, 1, 2, 2},
					{2, 0, 2, 2, 0},
					{2, 0, 1, 1, 0},
					{1, 0, 2, 2, 2},
					{2, 0, 0, 2, 2},
				},
			},
			ans3459: ans3459{
				result: 5,
			},
		},
		{
			name: "example 2",
			para3459: para3459{
				grid: [][]int{
					{2, 2, 2, 2, 2},
					{2, 0, 2, 2, 0},
					{2, 0, 1, 1, 0},
					{1, 0, 2, 2, 2},
					{2, 0, 0, 2, 2},
				},
			},
			ans3459: ans3459{
				result: 4,
			},
		},
		{
			name: "example 3",
			para3459: para3459{
				grid: [][]int{
					{1, 2, 2, 2, 2},
					{2, 2, 2, 2, 0},
					{2, 0, 0, 0, 0},
					{0, 0, 2, 2, 2},
					{2, 0, 0, 2, 0},
				},
			},
			ans3459: ans3459{
				result: 5,
			},
		},
		{
			name: "example 4",
			para3459: para3459{
				grid: [][]int{
					{1},
				},
			},
			ans3459: ans3459{
				result: 1,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3459------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3459, q.para3459
			output := lenOfVDiagonal(p.grid)
			fmt.Printf("[input]: grid=%v       [output]:%v\n", p.grid, output)
			assert.Equal(t, q.ans3459.result, output, "lenOfVDiagonal(%v) should return %v", p.grid, q.ans3459.result)
		})
	}
	fmt.Printf("\n\n\n")
}
