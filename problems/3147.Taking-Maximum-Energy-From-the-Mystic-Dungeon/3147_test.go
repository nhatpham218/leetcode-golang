package leetcode

import (
	"fmt"
	"testing"
)

type question3147 struct {
	name string
	para3147
	ans3147
}

type para3147 struct {
	energy []int
	k      int
}

type ans3147 struct {
	one int
}

func Test_Problem3147(t *testing.T) {

	qs := []question3147{
		{
			name: "example 1",
			para3147: para3147{
				energy: []int{5, 2, -10, -5, 1},
				k:      3,
			},
			ans3147: ans3147{3},
		},
		{
			name: "example 2",
			para3147: para3147{
				energy: []int{-2, -3, -1},
				k:      2,
			},
			ans3147: ans3147{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3147------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3147, q.para3147
			output := maximumEnergy(p.energy, p.k)
			fmt.Printf("[input]: energy=%v k=%v       [output]:%v\n", p.energy, p.k, output)
			if output != q.ans3147.one {
				t.Errorf("Expected %v, got %v", q.ans3147.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
