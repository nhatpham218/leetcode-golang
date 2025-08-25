package leetcode

import (
	"fmt"
	"testing"
)

type question3212 struct {
	para3212
	ans3212
}

// para is parameter
// one represents the first parameter
type para3212 struct {
	grid [][]byte
}

// ans is answer
// one represents the first answer
type ans3212 struct {
	one int
}

func Test_Problem3212(t *testing.T) {

	qs := []question3212{
		{
			para3212{[][]byte{
				{'X', 'Y', '.'},
				{'Y', '.', '.'},
			}},
			ans3212{3},
		},

		{
			para3212{[][]byte{
				{'X', 'X'},
				{'X', 'Y'},
			}},
			ans3212{0},
		},

		{
			para3212{[][]byte{
				{'.', '.'},
				{'.', '.'},
			}},
			ans3212{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3212------------------------\n")

	for i, q := range qs {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			_, p := q.ans3212, q.para3212
			output := numberOfSubmatrices(p.grid)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			if output != q.ans3212.one {
				t.Errorf("Expected %d, got %d", q.ans3212.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
