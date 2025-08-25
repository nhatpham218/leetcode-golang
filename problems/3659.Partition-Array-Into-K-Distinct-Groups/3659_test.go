package leetcode

import (
	"testing"
)

type question3659 struct {
	name string
	para3659
	ans3659
}

// para is parameter
// one represents the first parameter
type para3659 struct {
	one []int
	two int
}

// ans is answer
// one represents the first answer
type ans3659 struct {
	one bool
}

func Test_Problem3659(t *testing.T) {

	qs := []question3659{
		{
			name:     "example 1",
			para3659: para3659{[]int{1, 2, 3, 4}, 2},
			ans3659:  ans3659{true},
		},
		{
			name:     "example 2",
			para3659: para3659{[]int{3, 5, 2, 2}, 2},
			ans3659:  ans3659{true},
		},
		{
			name:     "example 3",
			para3659: para3659{[]int{1, 5, 2, 3}, 3},
			ans3659:  ans3659{false},
		},
		{
			name:     "single element, k=1",
			para3659: para3659{[]int{1}, 1},
			ans3659:  ans3659{true},
		},
		{
			name:     "two elements, k=2",
			para3659: para3659{[]int{1, 2}, 2},
			ans3659:  ans3659{true},
		},
		{
			name:     "two elements, k=1",
			para3659: para3659{[]int{1, 2}, 1},
			ans3659:  ans3659{true},
		},
	}

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3659, q.para3659
			output := partitionArray(p.one, p.two)
			if output != q.ans3659.one {
				t.Errorf("Case %s failed: expected %v, got %v", q.name, q.ans3659.one, output)
			}
		})
	}
}
