package leetcode

import (
	"fmt"
	"testing"
)

type question1488 struct {
	name string
	para1488
	ans1488
}

type para1488 struct {
	rains []int
}

type ans1488 struct {
	one []int
}

func Test_Problem1488(t *testing.T) {

	qs := []question1488{
		{
			name: "example 1",
			para1488: para1488{
				rains: []int{1, 2, 3, 4},
			},
			ans1488: ans1488{
				[]int{-1, -1, -1, -1},
			},
		},
		{
			name: "example 2",
			para1488: para1488{
				rains: []int{1, 2, 0, 0, 2, 1},
			},
			ans1488: ans1488{
				[]int{-1, -1, 2, 1, -1, -1},
			},
		},
		{
			name: "example 3",
			para1488: para1488{
				rains: []int{1, 2, 0, 1, 2},
			},
			ans1488: ans1488{
				[]int{},
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1488------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1488, q.para1488

			output := avoidFlood(p.rains)
			fmt.Printf("[input]: rains=%v       [output]:%v\n", p.rains, output)
			if !isValidSol(p.rains, output) {
				t.Errorf("Invalid output %v", output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

func isValidSol(rains []int, ans []int) bool {
	if len(ans) == 0 {
		return true
	}
	if len(rains) != len(ans) {
		return false
	}

	fullLakes := make(map[int]bool)

	for i := 0; i < len(rains); i++ {
		if rains[i] > 0 {
			if ans[i] != -1 {
				return false
			}
			if fullLakes[rains[i]] {
				return false
			}
			fullLakes[rains[i]] = true
		} else {
			if ans[i] <= 0 {
				return false
			}
			delete(fullLakes, ans[i])
		}
	}

	return true
}
