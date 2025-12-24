package leetcode

import (
	"fmt"
	"testing"
)

type question3074 struct {
	name string
	para3074
	ans3074
}

type para3074 struct {
	apple    []int
	capacity []int
}

type ans3074 struct {
	one int
}

func Test_Problem3074(t *testing.T) {

	qs := []question3074{
		{
			name: "example 1",
			para3074: para3074{
				apple:    []int{1, 3, 2},
				capacity: []int{4, 3, 1, 5, 2},
			},
			ans3074: ans3074{2},
		},
		{
			name: "example 2",
			para3074: para3074{
				apple:    []int{5, 5, 5},
				capacity: []int{2, 4, 2, 7},
			},
			ans3074: ans3074{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3074------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3074, q.para3074
			output := minimumBoxes(p.apple, p.capacity)
			fmt.Printf("[input]: apple=%v, capacity=%v       [output]:%v\n", p.apple, p.capacity, output)
			if output != q.ans3074.one {
				t.Errorf("Expected %v, got %v", q.ans3074.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

