package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3644 struct {
	name string
	para3644
	ans3644
}

// para is parameter
type para3644 struct {
	nums []int
}

// ans is answer
type ans3644 struct {
	k int
}

func Test_Problem3644(t *testing.T) {

	qs := []question3644{
		{
			name: "example 1",
			para3644: para3644{
				nums: []int{0, 3, 2, 1},
			},
			ans3644: ans3644{
				k: 1,
			},
		},
		{
			name: "example 2",
			para3644: para3644{
				nums: []int{0, 1, 3, 2},
			},
			ans3644: ans3644{
				k: 2,
			},
		},
		{
			name: "example 3",
			para3644: para3644{
				nums: []int{3, 2, 1, 0},
			},
			ans3644: ans3644{
				k: 0,
			},
		},
		{
			name: "already sorted",
			para3644: para3644{
				nums: []int{0, 1, 2, 3},
			},
			ans3644: ans3644{
				k: 0,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3644------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3644, q.para3644
			// Make a deep copy of the nums for testing since the function might modify the input
			numsCopy := make([]int, len(p.nums))
			copy(numsCopy, p.nums)
			output := sortPermutation(numsCopy)
			fmt.Printf("[input]: nums=%v       [output]:%d\n", p.nums, output)
			assert.Equal(t, q.ans3644.k, output, "sortPermutation(%v) should return %d", p.nums, q.ans3644.k)
		})
	}
	fmt.Printf("\n\n\n")
}
