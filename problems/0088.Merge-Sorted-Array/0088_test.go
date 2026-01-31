package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question88 struct {
	name string
	para88
	ans88
}

// para is parameter
// one represents the first parameter
type para88 struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

// ans is answer
// one represents the first answer
type ans88 struct {
	one []int
}

func Test_Problem88(t *testing.T) {

	qs := []question88{
		{
			name: "example 1",
			para88: para88{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			ans88: ans88{[]int{1, 2, 2, 3, 5, 6}},
		},
		{
			name: "example 2",
			para88: para88{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			ans88: ans88{[]int{1}},
		},
		{
			name: "example 3",
			para88: para88{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			ans88: ans88{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 88------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans88, q.para88
			merge(p.nums1, p.m, p.nums2, p.n)
			fmt.Printf("[input]: nums1=%v, m=%v, nums2=%v, n=%v       [output]:%v\n", p.nums1, p.m, p.nums2, p.n, p.nums1)
			if !reflect.DeepEqual(p.nums1, q.ans88.one) {
				t.Errorf("Expected %v, got %v", q.ans88.one, p.nums1)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
