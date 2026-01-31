package leetcode

import (
	"fmt"
	"testing"
)

type question80 struct {
	name string
	para80
	ans80
}

type para80 struct {
	nums []int
}

type ans80 struct {
	k       int
	resNums []int
}

func Test_Problem80(t *testing.T) {

	qs := []question80{
		{
			name:   "example 1",
			para80: para80{nums: []int{1, 1, 1, 2, 2, 3}},
			ans80:  ans80{k: 5, resNums: []int{1, 1, 2, 2, 3}},
		},
		{
			name:   "example 2",
			para80: para80{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			ans80:  ans80{k: 7, resNums: []int{0, 0, 1, 1, 2, 3, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 80------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			nums := make([]int, len(q.para80.nums))
			copy(nums, q.para80.nums)
			k := removeDuplicates(nums)
			fmt.Printf("[input]: nums=%v       [output]: k=%v, nums=%v\n", q.para80.nums, k, nums[:k])
			if k != q.ans80.k {
				t.Errorf("Expected k=%v, got k=%v", q.ans80.k, k)
			}
			for i := 0; i < k; i++ {
				if nums[i] != q.ans80.resNums[i] {
					t.Errorf("Expected nums[%d]=%v, got nums[%d]=%v", i, q.ans80.resNums[i], i, nums[i])
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
