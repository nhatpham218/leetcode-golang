package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question3379 struct {
	name string
	para3379
	ans3379
}

type para3379 struct {
	nums []int
}

type ans3379 struct {
	one []int
}

func Test_Problem3379(t *testing.T) {

	qs := []question3379{
		{
			name:     "example 1",
			para3379: para3379{nums: []int{3, -2, 1, 1}},
			ans3379:  ans3379{one: []int{1, 1, 1, 3}},
		},
		{
			name:     "example 2",
			para3379: para3379{nums: []int{-1, 4, -1}},
			ans3379:  ans3379{one: []int{-1, -1, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3379------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3379, q.para3379
			output := constructTransformedArray(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if !reflect.DeepEqual(output, q.ans3379.one) {
				t.Errorf("Expected %v, got %v", q.ans3379.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
