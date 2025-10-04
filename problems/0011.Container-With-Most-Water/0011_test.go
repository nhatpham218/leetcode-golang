package leetcode

import (
	"fmt"
	"testing"
)

type question11 struct {
	name string
	para11
	ans11
}

type para11 struct {
	height []int
}

type ans11 struct {
	one int
}

func Test_Problem11(t *testing.T) {

	qs := []question11{
		{
			name: "example 1",
			para11: para11{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			ans11: ans11{49},
		},
		{
			name: "example 2",
			para11: para11{
				height: []int{1, 1},
			},
			ans11: ans11{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 11------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans11, q.para11
			output := maxArea(p.height)
			fmt.Printf("[input]: height=%v       [output]:%d\n", p.height, output)
			if output != q.ans11.one {
				t.Errorf("Expected %d, got %d", q.ans11.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
