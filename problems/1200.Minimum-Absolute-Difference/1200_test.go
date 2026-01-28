package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question1200 struct {
	name string
	para1200
	ans1200
}

type para1200 struct {
	arr []int
}

type ans1200 struct {
	one [][]int
}

func Test_Problem1200(t *testing.T) {

	qs := []question1200{
		{
			name:     "example 1",
			para1200: para1200{arr: []int{4, 2, 1, 3}},
			ans1200:  ans1200{one: [][]int{{1, 2}, {2, 3}, {3, 4}}},
		},
		{
			name:     "example 2",
			para1200: para1200{arr: []int{1, 3, 6, 10, 15}},
			ans1200:  ans1200{one: [][]int{{1, 3}}},
		},
		{
			name:     "example 3",
			para1200: para1200{arr: []int{3, 8, -10, 23, 19, -4, -14, 27}},
			ans1200:  ans1200{one: [][]int{{-14, -10}, {19, 23}, {23, 27}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1200------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1200, q.para1200
			output := minimumAbsDifference(p.arr)
			fmt.Printf("[input]: arr=%v       [output]:%v\n", p.arr, output)
			if !reflect.DeepEqual(output, q.ans1200.one) {
				t.Errorf("Expected %v, got %v", q.ans1200.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
