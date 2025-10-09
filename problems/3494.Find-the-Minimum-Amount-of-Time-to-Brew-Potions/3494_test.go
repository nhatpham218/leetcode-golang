package leetcode

import (
	"fmt"
	"testing"
)

type question3494 struct {
	name string
	para3494
	ans3494
}

type para3494 struct {
	skill []int
	mana  []int
}

type ans3494 struct {
	one int64
}

func Test_Problem3494(t *testing.T) {

	qs := []question3494{
		{
			name: "example 1",
			para3494: para3494{
				skill: []int{1, 5, 2, 4},
				mana:  []int{5, 1, 4, 2},
			},
			ans3494: ans3494{110},
		},
		{
			name: "example 2",
			para3494: para3494{
				skill: []int{1, 1, 1},
				mana:  []int{1, 1, 1},
			},
			ans3494: ans3494{5},
		},
		{
			name: "example 3",
			para3494: para3494{
				skill: []int{1, 2, 3, 4},
				mana:  []int{1, 2},
			},
			ans3494: ans3494{21},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3494------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3494, q.para3494
			output := minTime(p.skill, p.mana)
			fmt.Printf("[input]: skill=%v mana=%v       [output]:%v\n", p.skill, p.mana, output)
			if output != q.ans3494.one {
				t.Errorf("Expected %v, got %v", q.ans3494.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
