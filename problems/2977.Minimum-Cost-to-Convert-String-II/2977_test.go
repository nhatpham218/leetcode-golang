package leetcode

import (
	"fmt"
	"testing"
)

type question2977 struct {
	name string
	para2977
	ans2977
}

// para is parameter
// one represents the first parameter
type para2977 struct {
	source   string
	target   string
	original []string
	changed  []string
	cost     []int
}

// ans is answer
// one represents the first answer
type ans2977 struct {
	one int64
}

func Test_Problem2977(t *testing.T) {

	qs := []question2977{
		{
			name: "example 1",
			para2977: para2977{
				source:   "abcd",
				target:   "acbe",
				original: []string{"a", "b", "c", "c", "e", "d"},
				changed:  []string{"b", "c", "b", "e", "b", "e"},
				cost:     []int{2, 5, 5, 1, 2, 20},
			},
			ans2977: ans2977{28},
		},
		{
			name: "example 2",
			para2977: para2977{
				source:   "abcdefgh",
				target:   "acdeeghh",
				original: []string{"bcd", "fgh", "thh"},
				changed:  []string{"cde", "thh", "ghh"},
				cost:     []int{1, 3, 5},
			},
			ans2977: ans2977{9},
		},
		{
			name: "example 3",
			para2977: para2977{
				source:   "abcdefgh",
				target:   "addddddd",
				original: []string{"bcd", "defgh"},
				changed:  []string{"ddd", "ddddd"},
				cost:     []int{100, 1578},
			},
			ans2977: ans2977{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2977------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2977, q.para2977
			output := minimumCost(p.source, p.target, p.original, p.changed, p.cost)
			fmt.Printf("[input]: source=%v, target=%v, original=%v, changed=%v, cost=%v       [output]:%v\n", p.source, p.target, p.original, p.changed, p.cost, output)
			if output != q.ans2977.one {
				t.Errorf("Expected %v, got %v", q.ans2977.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
