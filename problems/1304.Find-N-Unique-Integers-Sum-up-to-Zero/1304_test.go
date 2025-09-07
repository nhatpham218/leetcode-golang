package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question1304 struct {
	name string
	para1304
	ans1304
}

// para is parameter
// one represents the first parameter
type para1304 struct {
	one int
}

// ans is answer
// one represents the first answer
type ans1304 struct {
	one []int
}

func Test_Problem1304(t *testing.T) {

	qs := []question1304{
		{
			name:     "example 1",
			para1304: para1304{5},
			ans1304:  ans1304{[]int{-2, -1, 0, 1, 2}},
		},
		{
			name:     "example 2",
			para1304: para1304{3},
			ans1304:  ans1304{[]int{-1, 0, 1}},
		},
		{
			name:     "example 3",
			para1304: para1304{1},
			ans1304:  ans1304{[]int{0}},
		},
		{
			name:     "example 4",
			para1304: para1304{4},
			ans1304:  ans1304{[]int{-3, -2, -1, 0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1304------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1304, q.para1304
			output := sumZero(p.one)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)

			// Check if output has correct length
			assert.Equal(t, p.one, len(output), "sumZero(%v) should return array of length %v, got length %v", p.one, p.one, len(output))

			// Check if sum is zero
			sum := 0
			for _, val := range output {
				sum += val
			}
			assert.Equal(t, 0, sum, "sumZero(%v) should return array that sums to 0, got sum %v", p.one, sum)

			// Check if all elements are unique
			seen := make(map[int]bool)
			for _, val := range output {
				assert.False(t, seen[val], "sumZero(%v) should return array with unique elements, found duplicate %v", p.one, val)
				seen[val] = true
			}
		})
	}
	fmt.Printf("\n\n\n")
}
