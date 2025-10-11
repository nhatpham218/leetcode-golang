package leetcode

import (
	"fmt"
	"testing"
)

type question3186 struct {
	name string
	para3186
	ans3186
}

type para3186 struct {
	power []int
}

type ans3186 struct {
	one int64
}

func Test_Problem3186(t *testing.T) {

	qs := []question3186{
		{
			name: "example 1",
			para3186: para3186{
				power: []int{1, 1, 3, 4},
			},
			ans3186: ans3186{6},
		},
		{
			name: "example 2",
			para3186: para3186{
				power: []int{7, 1, 6, 6},
			},
			ans3186: ans3186{13},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3186------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3186, q.para3186
			output := maximumTotalDamage(p.power)
			fmt.Printf("[input]: power=%v       [output]:%v\n", p.power, output)
			if output != q.ans3186.one {
				t.Errorf("Expected %v, got %v", q.ans3186.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
