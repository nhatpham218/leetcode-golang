package leetcode

import (
	"fmt"
	"testing"
)

type question1733 struct {
	name string
	para1733
	ans1733
}

// para is parameter
// one represents the first parameter
type para1733 struct {
	n           int
	languages   [][]int
	friendships [][]int
}

// ans is answer
// one represents the first answer
type ans1733 struct {
	one int
}

func Test_Problem1733(t *testing.T) {

	qs := []question1733{
		{
			name: "example 1",
			para1733: para1733{
				n:           2,
				languages:   [][]int{{1}, {2}, {1, 2}},
				friendships: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			ans1733: ans1733{1},
		},
		{
			name: "example 2",
			para1733: para1733{
				n:           3,
				languages:   [][]int{{2}, {1, 3}, {1, 2}, {3}},
				friendships: [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}},
			},
			ans1733: ans1733{2},
		},
		{
			name: "all can communicate",
			para1733: para1733{
				n:           2,
				languages:   [][]int{{1, 2}, {1, 2}},
				friendships: [][]int{{1, 2}},
			},
			ans1733: ans1733{0},
		},
		{
			name: "single language optimal",
			para1733: para1733{
				n:           3,
				languages:   [][]int{{1}, {2}, {3}, {1}},
				friendships: [][]int{{1, 2}, {2, 3}, {3, 4}},
			},
			ans1733: ans1733{2},
		},
		{
			name: "edge case - two users",
			para1733: para1733{
				n:           2,
				languages:   [][]int{{1}, {2}},
				friendships: [][]int{{1, 2}},
			},
			ans1733: ans1733{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1733------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1733, q.para1733
			output := minimumTeachings(p.n, p.languages, p.friendships)
			fmt.Printf("[input]: n=%v, languages=%v, friendships=%v       [output]:%v\n", p.n, p.languages, p.friendships, output)
			if output != q.ans1733.one {
				t.Errorf("Expected %v, got %v", q.ans1733.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
