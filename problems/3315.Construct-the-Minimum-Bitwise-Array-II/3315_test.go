package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question3315 struct {
	name string
	para3315
	ans3315
}

// para is parameter
// one represents the first parameter
type para3315 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans3315 struct {
	one []int
}

func Test_Problem3315(t *testing.T) {

	qs := []question3315{
		{
			name: "example 1",
			para3315: para3315{
				nums: []int{2, 3, 5, 7},
			},
			ans3315: ans3315{[]int{-1, 1, 4, 3}},
		},
		{
			name: "example 2",
			para3315: para3315{
				nums: []int{11, 13, 31},
			},
			ans3315: ans3315{[]int{9, 12, 15}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3315------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3315, q.para3315
			output := minBitwiseArray(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if !reflect.DeepEqual(output, q.ans3315.one) {
				t.Errorf("Expected %v, got %v", q.ans3315.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
