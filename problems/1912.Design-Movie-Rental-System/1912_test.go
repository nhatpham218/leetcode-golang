package leetcode

import (
	"fmt"
	"testing"
)

type question1912 struct {
	name string
	para1912
	ans1912
}

// para is parameter
type para1912 struct {
	n       int
	entries [][]int
	ops     []string
	params  []interface{}
}

// ans is answer
type ans1912 struct {
	expected []interface{}
}

func Test_Problem1912(t *testing.T) {

	qs := []question1912{
		{
			name: "example 1",
			para1912: para1912{
				n:       3,
				entries: [][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}},
				ops:     []string{"search", "rent", "rent", "report", "drop", "search"},
				params: []interface{}{
					[]interface{}{1},
					[]interface{}{0, 1},
					[]interface{}{1, 2},
					[]interface{}{},
					[]interface{}{1, 2},
					[]interface{}{2},
				},
			},
			ans1912: ans1912{[]interface{}{[]int{1, 0, 2}, nil, nil, [][]int{{0, 1}, {1, 2}}, nil, []int{0, 1}}},
		},
		{
			name: "failing case",
			para1912: para1912{
				n:       10,
				entries: [][]int{{4, 374, 55}, {1, 6371, 21}, {8, 3660, 24}, {1, 56, 32}, {5, 374, 71}, {3, 4408, 36}, {6, 9322, 73}, {6, 9574, 92}, {8, 7834, 62}, {2, 6084, 27}, {7, 3262, 89}, {2, 8959, 53}, {0, 3323, 41}, {6, 6565, 45}, {0, 4239, 20}},
				ops:     []string{"rent", "drop", "rent", "rent", "rent", "drop", "search", "report", "rent", "search"},
				params: []interface{}{
					[]interface{}{0, 4239},
					[]interface{}{0, 4239},
					[]interface{}{3, 4408},
					[]interface{}{2, 6084},
					[]interface{}{0, 4239},
					[]interface{}{0, 4239},
					[]interface{}{9346},
					[]interface{}{},
					[]interface{}{6, 9322},
					[]interface{}{8698},
				},
			},
			ans1912: ans1912{[]interface{}{nil, nil, nil, nil, nil, nil, []int{}, [][]int{{2, 6084}, {3, 4408}}, nil, []int{}}},
		},
		{
			name: "complex case",
			para1912: para1912{
				n:       69,
				entries: [][]int{{16, 4156, 1511}, {20, 8501, 8417}, {34, 7901, 7776}, {54, 6691, 9511}, {44, 8931, 8434}, {42, 9640, 5251}, {22, 4534, 9161}, {32, 6506, 6831}, {13, 8501, 731}, {4, 7610, 8474}, {33, 820, 2341}, {17, 6490, 1161}, {29, 7120, 2703}, {8, 8723, 7613}, {38, 9544, 1804}, {30, 8723, 1047}, {1, 5015, 7763}, {60, 1625, 2383}, {29, 3336, 3542}, {39, 7535, 6066}, {1, 9074, 9400}, {39, 1625, 7944}, {26, 9160, 6874}, {55, 2465, 888}, {35, 8530, 6025}},
				ops:     []string{"rent", "search", "search", "report", "rent", "rent", "report", "report", "search", "search", "rent", "rent", "search", "drop", "drop", "drop", "drop", "rent", "report", "report", "rent", "drop", "search", "report", "drop", "report", "drop", "rent", "report", "search", "search", "rent", "rent", "report", "report", "drop", "report", "report", "drop", "report", "drop", "rent", "drop", "search", "rent", "search", "drop", "rent", "drop", "report", "rent", "drop", "rent", "rent", "drop", "report", "report", "report", "report", "rent", "drop", "report", "drop", "rent", "search", "drop", "report", "rent", "search", "search", "report", "rent", "report", "report", "rent", "report", "report", "search", "rent", "rent", "search"},
				params: []interface{}{
					[]interface{}{32, 6506}, []interface{}{8501}, []interface{}{6275}, []interface{}{},
					[]interface{}{30, 8723}, []interface{}{8, 8723}, []interface{}{}, []interface{}{},
					[]interface{}{6699}, []interface{}{115}, []interface{}{20, 8501}, []interface{}{16, 4156},
					[]interface{}{9447}, []interface{}{30, 8723}, []interface{}{8, 8723}, []interface{}{32, 6506},
					[]interface{}{16, 4156}, []interface{}{42, 9640}, []interface{}{}, []interface{}{},
					[]interface{}{17, 6490}, []interface{}{20, 8501}, []interface{}{8175}, []interface{}{},
					[]interface{}{17, 6490}, []interface{}{}, []interface{}{42, 9640}, []interface{}{54, 6691},
					[]interface{}{}, []interface{}{1625}, []interface{}{3291}, []interface{}{60, 1625},
					[]interface{}{39, 1625}, []interface{}{}, []interface{}{}, []interface{}{60, 1625},
					[]interface{}{}, []interface{}{}, []interface{}{39, 1625}, []interface{}{},
					[]interface{}{54, 6691}, []interface{}{8, 8723}, []interface{}{8, 8723}, []interface{}{2260},
					[]interface{}{29, 7120}, []interface{}{746}, []interface{}{29, 7120}, []interface{}{38, 9544},
					[]interface{}{38, 9544}, []interface{}{}, []interface{}{1, 9074}, []interface{}{1, 9074},
					[]interface{}{54, 6691}, []interface{}{39, 1625}, []interface{}{54, 6691}, []interface{}{},
					[]interface{}{}, []interface{}{}, []interface{}{}, []interface{}{26, 9160},
					[]interface{}{26, 9160}, []interface{}{}, []interface{}{39, 1625}, []interface{}{42, 9640},
					[]interface{}{9640}, []interface{}{42, 9640}, []interface{}{}, []interface{}{29, 7120},
					[]interface{}{5630}, []interface{}{1842}, []interface{}{}, []interface{}{16, 4156},
					[]interface{}{}, []interface{}{}, []interface{}{1, 9074}, []interface{}{},
					[]interface{}{}, []interface{}{7992}, []interface{}{4, 7610}, []interface{}{29, 3336},
					[]interface{}{1333},
				},
			},
			ans1912: ans1912{[]interface{}{nil, []int{13, 20}, []int{}, [][]int{{32, 6506}}, nil, nil, [][]int{{30, 8723}, {32, 6506}, {8, 8723}}, [][]int{{30, 8723}, {32, 6506}, {8, 8723}}, []int{}, []int{}, nil, nil, []int{}, nil, nil, nil, nil, nil, [][]int{{42, 9640}, {20, 8501}}, [][]int{{42, 9640}, {20, 8501}}, nil, nil, []int{}, [][]int{{17, 6490}, {42, 9640}}, nil, [][]int{{42, 9640}}, nil, nil, [][]int{{54, 6691}}, []int{60, 39}, []int{}, nil, nil, [][]int{{60, 1625}, {39, 1625}, {54, 6691}}, [][]int{{60, 1625}, {39, 1625}, {54, 6691}}, nil, [][]int{{39, 1625}, {54, 6691}}, [][]int{{39, 1625}, {54, 6691}}, nil, [][]int{{54, 6691}}, nil, nil, nil, []int{}, nil, []int{}, nil, nil, nil, [][]int{}, nil, nil, nil, nil, nil, [][]int{{39, 1625}}, [][]int{{39, 1625}}, [][]int{{39, 1625}}, [][]int{{39, 1625}}, nil, nil, [][]int{{39, 1625}}, nil, nil, []int{}, nil, [][]int{}, nil, []int{}, []int{}, [][]int{{29, 7120}}, nil, [][]int{{16, 4156}, {29, 7120}}, [][]int{{16, 4156}, {29, 7120}}, nil, [][]int{{16, 4156}, {29, 7120}, {1, 9074}}, [][]int{{16, 4156}, {29, 7120}, {1, 9074}}, []int{}, nil, nil, []int{}}},
		},
		{
			name: "search issue case",
			para1912: para1912{
				n:       3,
				entries: [][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}},
				ops:     []string{"search", "rent", "rent", "report", "drop", "search"},
				params: []interface{}{
					[]interface{}{1},
					[]interface{}{0, 1},
					[]interface{}{1, 2},
					[]interface{}{},
					[]interface{}{1, 2},
					[]interface{}{2},
				},
			},
			ans1912: ans1912{[]interface{}{[]int{1, 0, 2}, nil, nil, [][]int{{0, 1}, {1, 2}}, nil, []int{0, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1912------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1912, q.para1912
			obj := Constructor(p.n, p.entries)

			fmt.Printf("[input]: n=%v, entries=%v\n", p.n, p.entries)

			for i, op := range p.ops {
				switch op {
				case "search":
					params := p.params[i].([]interface{})
					result := obj.Search(params[0].(int))
					fmt.Printf("[operation]: %s(%d) = %v\n", op, params[0].(int), result)
				case "rent":
					params := p.params[i].([]interface{})
					obj.Rent(params[0].(int), params[1].(int))
					fmt.Printf("[operation]: %s(%d, %d)\n", op, params[0].(int), params[1].(int))
				case "drop":
					params := p.params[i].([]interface{})
					obj.Drop(params[0].(int), params[1].(int))
					fmt.Printf("[operation]: %s(%d, %d)\n", op, params[0].(int), params[1].(int))
				case "report":
					result := obj.Report()
					fmt.Printf("[operation]: %s() = %v\n", op, result)
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
