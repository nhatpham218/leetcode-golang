package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3643 struct {
	name string
	para3643
	ans3643
}

// para is parameter
type para3643 struct {
	grid [][]int
	x    int
	y    int
	k    int
}

// ans is answer
type ans3643 struct {
	grid [][]int
}

func Test_Problem3643(t *testing.T) {

	qs := []question3643{
		{
			name: "example 1",
			para3643: para3643{
				grid: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
				x:    1,
				y:    0,
				k:    3,
			},
			ans3643: ans3643{
				grid: [][]int{{1, 2, 3, 4}, {13, 14, 15, 8}, {9, 10, 11, 12}, {5, 6, 7, 16}},
			},
		},
		{
			name: "example 2",
			para3643: para3643{
				grid: [][]int{{3, 4, 2, 3}, {2, 3, 4, 2}},
				x:    0,
				y:    2,
				k:    2,
			},
			ans3643: ans3643{
				grid: [][]int{{3, 4, 4, 2}, {2, 3, 2, 3}},
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3643------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3643, q.para3643
			// Make a deep copy of the grid for testing since the function modifies the input
			gridCopy := make([][]int, len(p.grid))
			for i := range p.grid {
				gridCopy[i] = make([]int, len(p.grid[i]))
				copy(gridCopy[i], p.grid[i])
			}
			output := reverseSubmatrix(gridCopy, p.x, p.y, p.k)
			fmt.Printf("[input]: grid=%v, x=%d, y=%d, k=%d       [output]:%v\n", p.grid, p.x, p.y, p.k, output)
			assert.Equal(t, q.ans3643.grid, output, "reverseSubmatrix(%v, %d, %d, %d) should return %v", p.grid, p.x, p.y, p.k, q.ans3643.grid)
		})
	}
	fmt.Printf("\n\n\n")
}
