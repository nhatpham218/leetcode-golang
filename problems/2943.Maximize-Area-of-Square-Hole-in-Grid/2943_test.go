package leetcode

import (
	"fmt"
	"testing"
)

type question2943 struct {
	name string
	para2943
	ans2943
}

// para is parameter
// one represents the first parameter
type para2943 struct {
	n     int
	m     int
	hBars []int
	vBars []int
}

// ans is answer
// one represents the first answer
type ans2943 struct {
	one int
}

func Test_Problem2943(t *testing.T) {

	qs := []question2943{
		{
			name: "example 1",
			para2943: para2943{
				n:     2,
				m:     1,
				hBars: []int{2, 3},
				vBars: []int{2},
			},
			ans2943: ans2943{4},
		},
		{
			name: "example 2",
			para2943: para2943{
				n:     1,
				m:     1,
				hBars: []int{2},
				vBars: []int{2},
			},
			ans2943: ans2943{4},
		},
		{
			name: "example 3",
			para2943: para2943{
				n:     2,
				m:     3,
				hBars: []int{2, 3},
				vBars: []int{2, 4},
			},
			ans2943: ans2943{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2943------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2943, q.para2943
			output := maximizeSquareHoleArea(p.n, p.m, p.hBars, p.vBars)
			fmt.Printf("[input]: n=%v, m=%v, hBars=%v, vBars=%v       [output]:%v\n", p.n, p.m, p.hBars, p.vBars, output)
			if output != q.ans2943.one {
				t.Errorf("Expected %v, got %v", q.ans2943.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
