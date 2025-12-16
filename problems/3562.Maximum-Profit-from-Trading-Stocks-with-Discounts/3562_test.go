package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3562 struct {
	name string
	para3562
	ans3562
}

type para3562 struct {
	n         int
	present   []int
	future    []int
	hierarchy [][]int
	budget    int
}

type ans3562 struct {
	result int
}

func Test_Problem3562(t *testing.T) {
	qs := []question3562{
		{
			name: "example 1",
			para3562: para3562{
				n:         2,
				present:   []int{1, 2},
				future:    []int{4, 3},
				hierarchy: [][]int{{1, 2}},
				budget:    3,
			},
			ans3562: ans3562{result: 5},
		},
		{
			name: "example 2",
			para3562: para3562{
				n:         2,
				present:   []int{3, 4},
				future:    []int{5, 8},
				hierarchy: [][]int{{1, 2}},
				budget:    4,
			},
			ans3562: ans3562{result: 4},
		},
		{
			name: "example 3",
			para3562: para3562{
				n:         3,
				present:   []int{4, 6, 8},
				future:    []int{7, 9, 11},
				hierarchy: [][]int{{1, 2}, {1, 3}},
				budget:    10,
			},
			ans3562: ans3562{result: 10},
		},
		{
			name: "example 4",
			para3562: para3562{
				n:         3,
				present:   []int{5, 2, 3},
				future:    []int{8, 5, 6},
				hierarchy: [][]int{{1, 2}, {2, 3}},
				budget:    7,
			},
			ans3562: ans3562{result: 12},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3562------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3562, q.para3562
			output := maxProfit(p.n, p.present, p.future, p.hierarchy, p.budget)
			fmt.Printf("[input]: n=%v, present=%v, future=%v, hierarchy=%v, budget=%v       [output]:%v\n",
				p.n, p.present, p.future, p.hierarchy, p.budget, output)
			assert.Equal(t, q.ans3562.result, output, "maxProfit should return %v, got %v", q.ans3562.result, output)
		})
	}
	fmt.Printf("\n\n\n")
}
