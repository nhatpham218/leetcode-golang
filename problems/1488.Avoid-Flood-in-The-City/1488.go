package leetcode

// 1488. Avoid Flood in The City
// https://leetcode.com/problems/avoid-flood-in-the-city/
func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	lastRain := make(map[int]int)
	dryDays := []int{}

	for i := 0; i < n; i++ {
		if rains[i] == 0 {
			dryDays = append(dryDays, i)
			ans[i] = 1
		} else {
			lake := rains[i]
			ans[i] = -1

			if prevDay, exists := lastRain[lake]; exists {
				idx := binarySearch(dryDays, prevDay)

				if idx == len(dryDays) || dryDays[idx] >= i {
					return []int{}
				}

				dryDay := dryDays[idx]
				ans[dryDay] = lake

				dryDays = append(dryDays[:idx], dryDays[idx+1:]...)
			}

			lastRain[lake] = i
		}
	}

	return ans
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
