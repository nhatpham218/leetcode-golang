package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question189 struct {
	name string
	para189
	ans189
}

type para189 struct {
	nums []int
	k    int
}

type ans189 struct {
	one []int
}

func Test_Problem189(t *testing.T) {

	qs := []question189{
		{
			name:    "example 1",
			para189: para189{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			ans189:  ans189{[]int{5, 6, 7, 1, 2, 3, 4}},
		},
		{
			name:    "example 2",
			para189: para189{[]int{-1, -100, 3, 99}, 2},
			ans189:  ans189{[]int{3, 99, -1, -100}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 189------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans189, q.para189
			nums := make([]int, len(p.nums))
			copy(nums, p.nums)
			rotate(nums, p.k)
			fmt.Printf("[input]: nums=%v, k=%v       [output]:%v\n", p.nums, p.k, nums)
			if !reflect.DeepEqual(nums, q.ans189.one) {
				t.Errorf("Expected %v, got %v", q.ans189.one, nums)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
