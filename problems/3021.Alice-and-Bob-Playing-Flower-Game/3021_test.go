package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3021 struct {
	name string
	para3021
	ans3021
}

type para3021 struct {
	n int
	m int
}

type ans3021 struct {
	result int64
}

func Test_Problem3021(t *testing.T) {
	qs := []question3021{
		{
			name:     "example 1",
			para3021: para3021{3, 2},
			ans3021:  ans3021{3},
		},
		{
			name:     "example 2",
			para3021: para3021{1, 1},
			ans3021:  ans3021{0},
		},
		{
			name:     "small case",
			para3021: para3021{2, 3},
			ans3021:  ans3021{3},
		},
		{
			name:     "larger case",
			para3021: para3021{10, 15},
			ans3021:  ans3021{75},
		},
		{
			name:     "edge case 1",
			para3021: para3021{1, 2},
			ans3021:  ans3021{1},
		},
		{
			name:     "edge case 2",
			para3021: para3021{2, 1},
			ans3021:  ans3021{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3021------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3021, q.para3021
			output := flowerGame(p.n, p.m)
			fmt.Printf("[input]: n=%v, m=%v       [output]:%v\n", p.n, p.m, output)
			assert.Equal(t, q.ans3021.result, output, "flowerGame(%v, %v) should return %v, got %v", p.n, p.m, q.ans3021.result, output)
		})
	}
	fmt.Printf("\n\n\n")
}
