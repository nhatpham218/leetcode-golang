package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question1718 struct {
	name string
	para1718
	ans1718
}

type para1718 struct {
	n int
}

type ans1718 struct {
	one []int
}

func Test_Problem1718(t *testing.T) {
	qs := []question1718{
		{
			name:     "example 1",
			para1718: para1718{n: 3},
			ans1718:  ans1718{one: []int{3, 1, 2, 3, 2}},
		},
		{
			name:     "example 2",
			para1718: para1718{n: 5},
			ans1718:  ans1718{one: []int{5, 3, 1, 4, 3, 5, 2, 4, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1718------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1718, q.para1718
			output := constructDistancedSequence(p.n)
			fmt.Printf("[input]: n=%v       [output]:%v\n", p.n, output)
			if !reflect.DeepEqual(output, q.ans1718.one) {
				t.Errorf("Expected %v, got %v", q.ans1718.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
