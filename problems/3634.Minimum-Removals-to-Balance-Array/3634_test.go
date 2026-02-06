package leetcode

import (
	"fmt"
	"testing"
)

type question3634 struct {
	name string
	para3634
	ans3634
}

type para3634 struct {
	nums []int
	k    int
}

type ans3634 struct {
	one int
}

func Test_Problem3634(t *testing.T) {

	qs := []question3634{
		{
			name:     "example 1",
			para3634: para3634{nums: []int{2, 1, 5}, k: 2},
			ans3634:  ans3634{1},
		},
		{
			name:     "example 2",
			para3634: para3634{nums: []int{1, 6, 2, 9}, k: 3},
			ans3634:  ans3634{2},
		},
		{
			name:     "example 3",
			para3634: para3634{nums: []int{4, 6}, k: 2},
			ans3634:  ans3634{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3634------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3634, q.para3634
			output := minRemoval(p.nums, p.k)
			fmt.Printf("[input]: nums=%v, k=%v       [output]:%v\n", p.nums, p.k, output)
			if output != q.ans3634.one {
				t.Errorf("Expected %v, got %v", q.ans3634.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
