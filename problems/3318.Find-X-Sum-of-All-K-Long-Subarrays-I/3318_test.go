package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question3318 struct {
	name string
	para3318
	ans3318
}

// para is parameter
// one represents the first parameter
type para3318 struct {
	nums []int
	k    int
	x    int
}

// ans is answer
// one represents the first answer
type ans3318 struct {
	one []int
}

func Test_Problem3318(t *testing.T) {

	qs := []question3318{
		{
			name: "example 1",
			para3318: para3318{
				nums: []int{1, 1, 2, 2, 3, 4, 2, 3},
				k:    6,
				x:    2,
			},
			ans3318: ans3318{[]int{6, 10, 12}},
		},
		{
			name: "example 2",
			para3318: para3318{
				nums: []int{3, 8, 7, 8, 7, 5},
				k:    2,
				x:    2,
			},
			ans3318: ans3318{[]int{11, 15, 15, 15, 12}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3318------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3318, q.para3318
			output := findXSum(p.nums, p.k, p.x)
			fmt.Printf("[input]: nums=%v k=%v x=%v       [output]:%v\n", p.nums, p.k, p.x, output)
			if !reflect.DeepEqual(output, q.ans3318.one) {
				t.Errorf("Expected %v, got %v", q.ans3318.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

