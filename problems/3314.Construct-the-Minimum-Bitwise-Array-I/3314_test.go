package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question3314 struct {
	name string
	para3314
	ans3314
}

type para3314 struct {
	nums []int
}

type ans3314 struct {
	one []int
}

func Test_Problem3314(t *testing.T) {

	qs := []question3314{
		{
			name:     "example 1",
			para3314: para3314{nums: []int{2, 3, 5, 7}},
			ans3314:  ans3314{one: []int{-1, 1, 4, 3}},
		},
		{
			name:     "example 2",
			para3314: para3314{nums: []int{11, 13, 31}},
			ans3314:  ans3314{one: []int{9, 12, 15}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3314------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3314, q.para3314
			output := minBitwiseArray(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if !reflect.DeepEqual(output, q.ans3314.one) {
				t.Errorf("Expected %v, got %v", q.ans3314.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
