package leetcode

// 3640. Trionic Array II
// https://leetcode.com/problems/trionic-array-ii/description/
func maxSumTrionic(nums []int) int64 {
	n := len(nums)
	const NEG = int64(-1e18)

	d0 := make([]int64, n)
	d1 := make([]int64, n)
	d2 := make([]int64, n)
	d3 := make([]int64, n)

	d0[0] = int64(nums[0])
	d1[0], d2[0], d3[0] = NEG, NEG, NEG

	for i := 1; i < n; i++ {
		d0[i], d1[i], d2[i], d3[i] = int64(nums[i]), NEG, NEG, NEG
		v := int64(nums[i])

		if nums[i] > nums[i-1] {
			d0[i] = max(v, d0[i-1]+v)
			d1[i] = max(d1[i-1]+v, d0[i-1]+v)
			d3[i] = max(d3[i-1]+v, d2[i-1]+v)
		} else if nums[i] < nums[i-1] {
			d2[i] = max(d2[i-1]+v, d1[i-1]+v)
		}
	}

	res := NEG
	for i := 0; i < n; i++ {
		res = max(res, d3[i])
	}
	return res
}
