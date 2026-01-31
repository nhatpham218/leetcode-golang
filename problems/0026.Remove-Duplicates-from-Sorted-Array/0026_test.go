package leetcode

import (
	"fmt"
	"testing"
)

type question26 struct {
	name string
	para26
	ans26
}

type para26 struct {
	nums []int
}

type ans26 struct {
	k        int
	expected []int
}

func Test_Problem26(t *testing.T) {

	qs := []question26{
		{
			name: "example 1",
			para26: para26{
				nums: []int{1, 1, 2},
			},
			ans26: ans26{
				k:        2,
				expected: []int{1, 2},
			},
		},
		{
			name: "example 2",
			para26: para26{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			ans26: ans26{
				k:        5,
				expected: []int{0, 1, 2, 3, 4},
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 26------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans26, q.para26
			nums := make([]int, len(p.nums))
			copy(nums, p.nums)
			output := removeDuplicates(nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans26.k {
				t.Errorf("Expected k=%v, got %v", q.ans26.k, output)
			}
			for i := 0; i < q.ans26.k; i++ {
				if nums[i] != q.ans26.expected[i] {
					t.Errorf("Expected nums[%d]=%v, got %v", i, q.ans26.expected[i], nums[i])
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
