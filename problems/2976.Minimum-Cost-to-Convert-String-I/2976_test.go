package leetcode

import (
	"fmt"
	"testing"
)

type question2976 struct {
	name string
	para2976
	ans2976
}

// para is parameter
// one represents the first parameter
type para2976 struct {
	source   string
	target   string
	original []byte
	changed  []byte
	cost     []int
}

// ans is answer
// one represents the first answer
type ans2976 struct {
	one int64
}

func Test_Problem2976(t *testing.T) {

	qs := []question2976{
		{
			name: "example 1",
			para2976: para2976{
				source:   "abcd",
				target:   "acbe",
				original: []byte{'a', 'b', 'c', 'c', 'e', 'd'},
				changed:  []byte{'b', 'c', 'b', 'e', 'b', 'e'},
				cost:     []int{2, 5, 5, 1, 2, 20},
			},
			ans2976: ans2976{28},
		},
		{
			name: "example 2",
			para2976: para2976{
				source:   "aaaa",
				target:   "bbbb",
				original: []byte{'a', 'c'},
				changed:  []byte{'c', 'b'},
				cost:     []int{1, 2},
			},
			ans2976: ans2976{12},
		},
		{
			name: "example 3",
			para2976: para2976{
				source:   "abcd",
				target:   "abce",
				original: []byte{'a'},
				changed:  []byte{'e'},
				cost:     []int{10000},
			},
			ans2976: ans2976{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2976------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2976, q.para2976
			output := minimumCost(p.source, p.target, p.original, p.changed, p.cost)
			fmt.Printf("[input]: source=%v, target=%v, original=%v, changed=%v, cost=%v       [output]:%v\n", p.source, p.target, p.original, p.changed, p.cost, output)
			if output != q.ans2976.one {
				t.Errorf("Expected %v, got %v", q.ans2976.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
