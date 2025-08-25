package leetcode

import (
	"fmt"
	"testing"
)

type question3661 struct {
	name string
	para3661
	ans3661
}

// para is parameter
// one represents the first parameter
type para3661 struct {
	robots   []int
	distance []int
	walls    []int
}

// ans is answer
// one represents the first answer
type ans3661 struct {
	one int
}

func Test_Problem3661(t *testing.T) {

	qs := []question3661{
		{
			name: "example 1",
			para3661: para3661{
				robots:   []int{4},
				distance: []int{3},
				walls:    []int{1, 10},
			},
			ans3661: ans3661{1},
		},
		{
			name: "example 2",
			para3661: para3661{
				robots:   []int{10, 2},
				distance: []int{5, 1},
				walls:    []int{5, 2, 7},
			},
			ans3661: ans3661{3},
		},
		{
			name: "example 3",
			para3661: para3661{
				robots:   []int{1, 2},
				distance: []int{100, 1},
				walls:    []int{10},
			},
			ans3661: ans3661{0},
		},
		{
			name: "example 4",
			para3661: para3661{
				robots:   []int{63, 56, 40, 45, 4, 9, 44, 69, 55, 26, 73, 15, 12, 60, 43, 39, 37, 74, 36, 34, 13, 23, 66, 14, 11, 42, 72, 3, 57, 10, 53, 8, 70, 17, 58, 61, 30, 32},
				distance: []int{8, 7, 4, 8, 9, 5, 2, 4, 5, 2, 6, 9, 5, 9, 5, 3, 7, 6, 9, 2, 8, 7, 4, 3, 5, 1, 7, 5, 1, 3, 5, 3, 5, 4, 8, 7, 6, 4},
				walls:    []int{6, 22, 50, 52, 20, 9, 23, 75, 26, 21, 60, 58, 41, 28, 30},
			},
			ans3661: ans3661{15},
		},
		{
			name: "example 5",
			para3661: para3661{
				robots:   []int{31, 22, 4, 43, 8, 38, 5, 15, 35, 37, 27, 42, 40, 28, 20, 21},
				distance: []int{3, 5, 5, 7, 8, 1, 10, 7, 9, 6, 3, 4, 4, 5, 7, 4},
				walls:    []int{34, 74, 54, 46, 79, 89, 7, 73, 12, 27, 44, 5, 62, 43, 60, 71, 10, 63, 41, 77, 33, 91, 32, 53, 66, 51, 78, 18, 61, 6, 8, 24, 23, 81, 3, 25, 40, 85, 84, 15, 52, 48, 17, 59, 55, 64, 50, 21, 88, 36, 2, 16, 80, 69, 22, 87, 1, 28, 65, 31, 83, 26, 67, 72, 29, 75, 57, 9, 30, 86, 39, 37, 13, 19, 56, 68, 35, 90},
			},
			ans3661: ans3661{41},
		},
		{
			name: "example 6",
			para3661: para3661{
				robots:   []int{47, 30, 21, 29, 56, 58, 49, 66, 39, 45, 34, 62, 40, 33, 51, 25, 7, 5, 52, 42, 61, 14, 69, 28, 54, 72, 9, 31, 59, 24, 18, 38, 10, 3, 8, 71},
				distance: []int{1, 1, 1, 2, 4, 4, 2, 5, 1, 1, 2, 5, 2, 5, 5, 4, 4, 3, 5, 3, 1, 4, 1, 1, 1, 3, 1, 2, 3, 4, 5, 2, 3, 5, 4, 2},
				walls:    []int{25, 45, 14, 42, 5, 22, 19, 35, 6, 44, 30, 10, 33, 11, 23, 21, 46, 49, 40, 8, 16, 20, 24, 41, 31, 9, 26, 47, 18, 4, 3, 32, 12, 27, 28, 38, 34, 1, 17, 39, 29, 13, 43, 7, 37, 48, 36, 2, 15},
			},
			ans3661: ans3661{49},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3661------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3661, q.para3661
			output := maxWalls(p.robots, p.distance, p.walls)
			fmt.Printf("[input]: robots=%v, distance=%v, walls=%v       [output]:%v\n", p.robots, p.distance, p.walls, output)
			if output != q.ans3661.one {
				t.Errorf("Expected %v, got %v", q.ans3661.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
