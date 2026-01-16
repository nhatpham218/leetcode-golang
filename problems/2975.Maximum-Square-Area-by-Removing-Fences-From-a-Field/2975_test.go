package leetcode

import (
	"fmt"
	"testing"
)

type question2975 struct {
	name string
	para2975
	ans2975
}

type para2975 struct {
	m       int
	n       int
	hFences []int
	vFences []int
}

type ans2975 struct {
	one int
}

func Test_Problem2975(t *testing.T) {

	qs := []question2975{
		{
			name: "example 1",
			para2975: para2975{
				m:       4,
				n:       3,
				hFences: []int{2, 3},
				vFences: []int{2},
			},
			ans2975: ans2975{4},
		},
		{
			name: "example 2",
			para2975: para2975{
				m:       6,
				n:       7,
				hFences: []int{2},
				vFences: []int{4},
			},
			ans2975: ans2975{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2975------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2975, q.para2975
			output := maximizeSquareArea(p.m, p.n, p.hFences, p.vFences)
			fmt.Printf("[input]: m=%v, n=%v, hFences=%v, vFences=%v       [output]:%v\n", p.m, p.n, p.hFences, p.vFences, output)
			if output != q.ans2975.one {
				t.Errorf("Expected %v, got %v", q.ans2975.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
