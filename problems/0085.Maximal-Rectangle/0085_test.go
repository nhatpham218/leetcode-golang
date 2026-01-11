package leetcode

import (
	"fmt"
	"testing"
)

type question0085 struct {
	name string
	para0085
	ans0085
}

// para is parameter
// one represents the first parameter
type para0085 struct {
	matrix [][]byte
}

// ans is answer
// one represents the first answer
type ans0085 struct {
	one int
}

func Test_Problem0085(t *testing.T) {

	qs := []question0085{
		{
			name: "example 1",
			para0085: para0085{
				matrix: [][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			ans0085: ans0085{6},
		},
		{
			name: "example 2",
			para0085: para0085{
				matrix: [][]byte{
					{'0'},
				},
			},
			ans0085: ans0085{0},
		},
		{
			name: "example 3",
			para0085: para0085{
				matrix: [][]byte{
					{'1'},
				},
			},
			ans0085: ans0085{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0085------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0085, q.para0085
			output := maximalRectangle(p.matrix)
			fmt.Printf("[input]: matrix=%v       [output]:%v\n", p.matrix, output)
			if output != q.ans0085.one {
				t.Errorf("Expected %v, got %v", q.ans0085.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
