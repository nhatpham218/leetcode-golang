package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

type question27 struct {
	name string
	para27
	ans27
}

type para27 struct {
	nums []int
	val  int
}

type ans27 struct {
	k        int
	expected []int
}

func Test_Problem27(t *testing.T) {

	qs := []question27{
		{
			name: "example 1",
			para27: para27{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			ans27: ans27{
				k:        2,
				expected: []int{2, 2},
			},
		},
		{
			name: "example 2",
			para27: para27{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			ans27: ans27{
				k:        5,
				expected: []int{0, 0, 1, 3, 4},
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 27------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			nums := make([]int, len(q.para27.nums))
			copy(nums, q.para27.nums)
			k := removeElement(nums, q.para27.val)
			fmt.Printf("[input]: nums=%v, val=%v       [output]: k=%v, nums=%v\n", q.para27.nums, q.para27.val, k, nums[:k])
			if k != q.ans27.k {
				t.Errorf("Expected k=%v, got k=%v", q.ans27.k, k)
			}
			sort.Ints(nums[:k])
			for i := 0; i < k; i++ {
				if nums[i] != q.ans27.expected[i] {
					t.Errorf("Expected nums[%d]=%v, got %v", i, q.ans27.expected[i], nums[i])
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
