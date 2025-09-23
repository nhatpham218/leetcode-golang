package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question165 struct {
	name string
	para165
	ans165
}

// para is parameter
type para165 struct {
	version1 string
	version2 string
}

// ans is answer
type ans165 struct {
	result int
}

func Test_Problem165(t *testing.T) {
	qs := []question165{
		{
			name:    "example 1",
			para165: para165{"1.2", "1.10"},
			ans165:  ans165{-1},
		},
		{
			name:    "example 2",
			para165: para165{"1.01", "1.001"},
			ans165:  ans165{0},
		},
		{
			name:    "example 3",
			para165: para165{"1.0", "1.0.0.0"},
			ans165:  ans165{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 165------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans165, q.para165
			output := compareVersion(p.version1, p.version2)
			fmt.Printf("[input]: version1=%s, version2=%s       [output]:%v\n", p.version1, p.version2, output)
			assert.Equal(t, q.ans165.result, output, "compareVersion(%s, %s) should return %v, got %v", p.version1, p.version2, q.ans165.result, output)
		})
	}
	fmt.Printf("\n\n\n")
}
