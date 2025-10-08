package leetcode

import (
	"fmt"
	"testing"
)

type question2300 struct {
	name string
	para2300
	ans2300
}

type para2300 struct {
	spells  []int
	potions []int
	success int64
}

type ans2300 struct {
	one []int
}

func Test_Problem2300(t *testing.T) {

	qs := []question2300{
		{
			name: "example 1",
			para2300: para2300{
				spells:  []int{5, 1, 3},
				potions: []int{1, 2, 3, 4, 5},
				success: 7,
			},
			ans2300: ans2300{[]int{4, 0, 3}},
		},
		{
			name: "example 2",
			para2300: para2300{
				spells:  []int{3, 1, 2},
				potions: []int{8, 5, 8},
				success: 16,
			},
			ans2300: ans2300{[]int{2, 0, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2300------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2300, q.para2300
			output := successfulPairs(p.spells, p.potions, p.success)
			fmt.Printf("[input]: spells=%v potions=%v success=%d       [output]:%v\n", p.spells, p.potions, p.success, output)
			if len(output) != len(q.ans2300.one) {
				t.Errorf("Expected %v, got %v", q.ans2300.one, output)
				return
			}
			for i := range output {
				if output[i] != q.ans2300.one[i] {
					t.Errorf("Expected %v, got %v", q.ans2300.one, output)
					return
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
