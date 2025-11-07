package leetcode

import (
	"fmt"
	"testing"
)

type question2528 struct {
	name string
	para2528
	ans2528
}

type para2528 struct {
	stations []int
	r        int
	k        int
}

type ans2528 struct {
	one int64
}

func Test_Problem2528(t *testing.T) {

	qs := []question2528{
		{
			name: "example 1",
			para2528: para2528{
				stations: []int{1, 2, 4, 5, 0},
				r:        1,
				k:        2,
			},
			ans2528: ans2528{5},
		},
		{
			name: "example 2",
			para2528: para2528{
				stations: []int{4, 4, 4, 4},
				r:        0,
				k:        3,
			},
			ans2528: ans2528{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2528------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2528, q.para2528
			output := maxPower(p.stations, p.r, p.k)
			fmt.Printf("[input]: stations=%v r=%v k=%v       [output]:%v\n", p.stations, p.r, p.k, output)
			if output != q.ans2528.one {
				t.Errorf("Expected %v, got %v", q.ans2528.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

