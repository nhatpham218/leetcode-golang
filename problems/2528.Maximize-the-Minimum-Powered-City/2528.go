package leetcode

// 2528. Maximize the Minimum Powered City
// https://leetcode.com/problems/maximize-the-minimum-powered-city/
func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)

	diffArr := make([]int64, n+1)
	for i := 0; i < n; i++ {
		left := max(0, i-r)
		right := min(n, i+r+1)
		diffArr[left] += int64(stations[i])
		diffArr[right] -= int64(stations[i])
	}

	minPower := int64(stations[0])
	totalPower := int64(0)
	for _, s := range stations {
		if int64(s) < minPower {
			minPower = int64(s)
		}
		totalPower += int64(s)
	}

	lo, hi := minPower, totalPower+int64(k)
	var result int64 = 0

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if canAchieve(diffArr, mid, r, k) {
			result = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func canAchieve(diffArr []int64, target int64, r int, k int) bool {
	n := len(diffArr) - 1
	tempDiff := make([]int64, len(diffArr))
	copy(tempDiff, diffArr)
	var power int64 = 0
	kLeft := int64(k)

	for i := 0; i < n; i++ {
		power += tempDiff[i]
		if power < target {
			needed := target - power
			if kLeft < needed {
				return false
			}
			kLeft -= needed
			end := min(n, i+2*r+1)
			tempDiff[end] -= needed
			power += needed
		}
	}

	return true
}
