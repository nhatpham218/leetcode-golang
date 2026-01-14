package leetcode

import "slices"

// 3454. Separate Squares II
// https://leetcode.com/problems/separate-squares-ii/description/

type Event struct {
	Y      int
	Type   int
	XStart int
	XEnd   int
}

type Area struct {
	Y      int
	Height float64
	Width  float64
}

type Interval struct {
	X1, X2 int
}

func getUnionWidth(intervals []Interval) float64 {
	if len(intervals) == 0 {
		return 0.0
	}
	slices.SortFunc(intervals, func(a, b Interval) int {
		if a.X1 != b.X1 {
			return a.X1 - b.X1
		}
		return a.X2 - b.X2
	})

	res := 0
	end := -2000000000
	for _, interval := range intervals {
		if interval.X1 > end {
			res += interval.X2 - interval.X1
			end = interval.X2
		} else if interval.X2 > end {
			res += interval.X2 - end
			end = interval.X2
		}
	}
	return float64(res)
}

func separateSquares(squares [][]int) float64 {
	var events []Event
	for _, sq := range squares {
		x, y, l := sq[0], sq[1], sq[2]
		events = append(events, Event{
			Y:      y,
			Type:   1,
			XStart: x,
			XEnd:   x + l,
		})
		events = append(events, Event{
			Y:      y + l,
			Type:   -1,
			XStart: x,
			XEnd:   x + l,
		})
	}
	slices.SortFunc(events, func(a, b Event) int {
		return a.Y - b.Y
	})

	var activeIntervals []Interval
	prevY := events[0].Y
	total := 0.0
	var areas []Area
	for _, event := range events {
		if event.Y > prevY {
			h := float64(event.Y - prevY)
			w := getUnionWidth(activeIntervals)
			areas = append(areas, Area{
				Y:      prevY,
				Height: h,
				Width:  w,
			})
			total += h * w
		}

		if event.Type == 1 {
			activeIntervals = append(activeIntervals, Interval{
				event.XStart,
				event.XEnd,
			})
		} else {
			idx := slices.IndexFunc(activeIntervals, func(inv Interval) bool {
				return inv.X1 == event.XStart && inv.X2 == event.XEnd
			})
			if idx != -1 {
				activeIntervals = append(activeIntervals[:idx], activeIntervals[idx+1:]...)
			}
		}
		prevY = event.Y
	}

	target := total / 2.0
	currentArea := 0.0
	for _, area := range areas {
		stripArea := area.Height * area.Width
		if currentArea+stripArea >= target {
			remaining := target - currentArea
			return float64(area.Y) + (remaining / area.Width)
		}
		currentArea += stripArea
	}

	return float64(events[len(events)-1].Y)
}
