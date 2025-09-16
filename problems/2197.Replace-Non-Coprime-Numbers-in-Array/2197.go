package leetcode

// 2197. Replace Non-Coprime Numbers in Array
// https://leetcode.com/problems/replace-non-coprime-numbers-in-array/

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func replaceNonCoprimes(nums []int) []int {
	stack := []int{}

	for _, num := range nums {
		stack = append(stack, num)

		// Keep merging with the previous element if they're not coprime
		for len(stack) >= 2 {
			top := stack[len(stack)-1]
			prev := stack[len(stack)-2]

			if gcd(top, prev) > 1 {
				// Remove both numbers and add their LCM
				stack = stack[:len(stack)-2]
				stack = append(stack, lcm(top, prev))
			} else {
				break
			}
		}
	}

	return stack
}
