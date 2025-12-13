package leetcode

import (
	"fmt"
	"testing"
)

type question3606 struct {
	name string
	para3606
	ans3606
}

type para3606 struct {
	code        []string
	businessLine []string
	isActive    []bool
}

type ans3606 struct {
	one []string
}

func Test_Problem3606(t *testing.T) {

	qs := []question3606{
		{
			name: "example 1",
			para3606: para3606{
				code:        []string{"SAVE20", "", "PHARMA5", "SAVE@20"},
				businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"},
				isActive:    []bool{true, true, true, true},
			},
			ans3606: ans3606{[]string{"PHARMA5", "SAVE20"}},
		},
		{
			name: "example 2",
			para3606: para3606{
				code:        []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"},
				businessLine: []string{"grocery", "electronics", "invalid"},
				isActive:    []bool{false, true, true},
			},
			ans3606: ans3606{[]string{"ELECTRONICS_50"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3606------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3606, q.para3606
			output := validateCoupons(p.code, p.businessLine, p.isActive)
			fmt.Printf("[input]: code=%v, businessLine=%v, isActive=%v       [output]:%v\n", p.code, p.businessLine, p.isActive, output)
			if len(output) != len(q.ans3606.one) {
				t.Errorf("Expected length %v, got %v", len(q.ans3606.one), len(output))
			}
			for i := range output {
				if output[i] != q.ans3606.one[i] {
					t.Errorf("Expected %v, got %v", q.ans3606.one, output)
					break
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}

