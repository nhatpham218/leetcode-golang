package leetcode

import (
	"fmt"
	"testing"
)

type question1504 struct {
	para1504
	ans1504
}

// para is parameter
// one represents the first parameter
type para1504 struct {
	mat [][]int
}

// ans is answer
// one represents the first answer
type ans1504 struct {
	one int
}

func Test_Problem1504(t *testing.T) {

	qs := []question1504{
		{
			para1504{[][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			}},
			ans1504{13},
		},

		{
			para1504{[][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 1},
				{1, 1, 1, 0},
			}},
			ans1504{24},
		},
		// If you need more tests, you can copy the elements above.
	}

	fmt.Printf("------------------------Leetcode Problem 1504------------------------\n")

	for i, q := range qs {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			_, p := q.ans1504, q.para1504
			output := numSubmat(p.mat)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			if output != q.ans1504.one {
				t.Errorf("Expected %d, got %d", q.ans1504.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
