package leetcode

import (
	"fmt"
	"testing"
)

type question2353 struct {
	name string
	para2353
	ans2353
}

// para is parameter
type para2353 struct {
	foods    []string
	cuisines []string
	ratings  []int
	ops      []string
	params   []interface{}
}

// ans is answer
type ans2353 struct {
	expected []interface{}
}

func Test_Problem2353(t *testing.T) {

	qs := []question2353{
		{
			name: "example 1",
			para2353: para2353{
				foods:    []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
				cuisines: []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
				ratings:  []int{9, 12, 8, 15, 14, 7},
				ops:      []string{"highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"},
				params:   []interface{}{"korean", "japanese", []interface{}{"sushi", 16}, "japanese", []interface{}{"ramen", 16}, "japanese"},
			},
			ans2353: ans2353{[]interface{}{"kimchi", "ramen", nil, "sushi", nil, "ramen"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2353------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2353, q.para2353
			obj := Constructor(p.foods, p.cuisines, p.ratings)

			fmt.Printf("[input]: foods=%v cuisines=%v ratings=%v\n", p.foods, p.cuisines, p.ratings)

			for i, op := range p.ops {
				switch op {
				case "highestRated":
					result := obj.HighestRated(p.params[i].(string))
					fmt.Printf("[operation]: %s(%s) = %s\n", op, p.params[i].(string), result)
				case "changeRating":
					params := p.params[i].([]interface{})
					obj.ChangeRating(params[0].(string), params[1].(int))
					fmt.Printf("[operation]: %s(%s, %d)\n", op, params[0].(string), params[1].(int))
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
