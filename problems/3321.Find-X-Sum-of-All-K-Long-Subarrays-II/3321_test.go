package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question3321 struct {
	name string
	para3321
	ans3321
}

// para is parameter
// one represents the first parameter
type para3321 struct {
	nums []int
	k    int
	x    int
}

// ans is answer
// one represents the first answer
type ans3321 struct {
	one []int64
}

func Test_Problem3321(t *testing.T) {

	qs := []question3321{
		{
			name: "example 1",
			para3321: para3321{
				nums: []int{1, 1, 2, 2, 3, 4, 2, 3},
				k:    6,
				x:    2,
			},
			ans3321: ans3321{[]int64{6, 10, 12}},
		},
		{
			name: "example 2",
			para3321: para3321{
				nums: []int{3, 8, 7, 8, 7, 5},
				k:    2,
				x:    2,
			},
			ans3321: ans3321{[]int64{11, 15, 15, 15, 12}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3321------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3321, q.para3321
			output := findXSum(p.nums, p.k, p.x)
			fmt.Printf("[input]: nums=%v k=%v x=%v       [output]:%v\n", p.nums, p.k, p.x, output)
			if !reflect.DeepEqual(output, q.ans3321.one) {
				t.Errorf("Expected %v, got %v", q.ans3321.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
