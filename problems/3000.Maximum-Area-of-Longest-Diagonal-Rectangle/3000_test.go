package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3000 struct {
	name string
	para3000
	ans3000
}

// para is parameter
// dimensions represents the first parameter
type para3000 struct {
	dimensions [][]int
}

// ans is answer
// one represents the first answer
type ans3000 struct {
	one int
}

func Test_Problem3000(t *testing.T) {

	qs := []question3000{
		{
			name:     "example 1",
			para3000: para3000{[][]int{{9, 3}, {8, 6}}},
			ans3000:  ans3000{48},
		},
		{
			name:     "example 2",
			para3000: para3000{[][]int{{3, 4}, {4, 3}}},
			ans3000:  ans3000{12},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3000------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3000, q.para3000
			output := areaOfMaxDiagonal(p.dimensions)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3000.one, output, "areaOfMaxDiagonal(%v) should return %d", p.dimensions, q.ans3000.one)
		})
	}
	fmt.Printf("\n\n\n")
}
