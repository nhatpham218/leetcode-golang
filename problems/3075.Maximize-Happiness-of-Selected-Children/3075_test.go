package leetcode

import (
	"fmt"
	"testing"
)

type question3075 struct {
	name string
	para3075
	ans3075
}

type para3075 struct {
	happiness []int
	k         int
}

type ans3075 struct {
	one int64
}

func Test_Problem3075(t *testing.T) {

	qs := []question3075{
		{
			name: "example 1",
			para3075: para3075{
				happiness: []int{1, 2, 3},
				k:         2,
			},
			ans3075: ans3075{4},
		},
		{
			name: "example 2",
			para3075: para3075{
				happiness: []int{1, 1, 1, 1},
				k:         2,
			},
			ans3075: ans3075{1},
		},
		{
			name: "example 3",
			para3075: para3075{
				happiness: []int{2, 3, 4, 5},
				k:         1,
			},
			ans3075: ans3075{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3075------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3075, q.para3075
			output := maximumHappinessSum(p.happiness, p.k)
			fmt.Printf("[input]: happiness=%v, k=%v       [output]:%v\n", p.happiness, p.k, output)
			if output != q.ans3075.one {
				t.Errorf("Expected %v, got %v", q.ans3075.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

