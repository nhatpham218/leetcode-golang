package leetcode

import (
	"fmt"
	"testing"
)

type question840 struct {
	name string
	para840
	ans840
}

type para840 struct {
	grid [][]int
}

type ans840 struct {
	one int
}

func Test_Problem840(t *testing.T) {

	qs := []question840{
		{
			name: "example 1",
			para840: para840{
				grid: [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}},
			},
			ans840: ans840{1},
		},
		{
			name: "example 2",
			para840: para840{
				grid: [][]int{{8}},
			},
			ans840: ans840{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 840------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans840, q.para840
			output := numMagicSquaresInside(p.grid)
			fmt.Printf("[input]: grid=%v       [output]:%v\n", p.grid, output)
			if output != q.ans840.one {
				t.Errorf("Expected %v, got %v", q.ans840.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

