package leetcode

import (
	"fmt"
	"testing"
)

type question1458 struct {
	name string
	para1458
	ans1458
}

// para is parameter
type para1458 struct {
	nums1 []int
	nums2 []int
}

// ans is answer
type ans1458 struct {
	one int
}

func Test_Problem1458(t *testing.T) {

	qs := []question1458{
		{
			name: "example 1",
			para1458: para1458{
				nums1: []int{2, 1, -2, 5},
				nums2: []int{3, 0, -6},
			},
			ans1458: ans1458{18},
		},
		{
			name: "example 2",
			para1458: para1458{
				nums1: []int{3, -2},
				nums2: []int{2, -6, 7},
			},
			ans1458: ans1458{21},
		},
		{
			name: "example 3",
			para1458: para1458{
				nums1: []int{-1, -1},
				nums2: []int{1, 1},
			},
			ans1458: ans1458{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1458------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1458, q.para1458
			output := maxDotProduct(p.nums1, p.nums2)
			fmt.Printf("[input]: nums1=%v, nums2=%v       [output]:%v\n", p.nums1, p.nums2, output)
			if output != q.ans1458.one {
				t.Errorf("Expected %v, got %v", q.ans1458.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

