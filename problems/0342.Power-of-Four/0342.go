package main

import "fmt"

// isPowerOfFour checks if n is a power of four
// A number is a power of four if:
// 1. It's positive
// 2. It's a power of 2 (only one bit set)
// 3. The set bit is at an even position (0, 2, 4, 6, ...)
func isPowerOfFour(n int) bool {
	// Check if n is positive and a power of 2
	if n <= 0 || n&(n-1) != 0 {
		return false
	}

	// Check if the set bit is at an even position
	// For powers of 4, the set bit should be at positions 0, 2, 4, 6, ...
	// We can check this by ensuring n-1 is divisible by 3
	// This works because 4^x - 1 = (4-1)(4^(x-1) + 4^(x-2) + ... + 1) = 3 * something
	return (n-1)%3 == 0
}

// Alternative solution using bit manipulation
func isPowerOfFourBitwise(n int) bool {
	// Check if n is positive and a power of 2
	if n <= 0 || n&(n-1) != 0 {
		return false
	}

	// Check if the set bit is at an even position
	// We can use a mask to check if the set bit is in the right position
	// 0x55555555 = 01010101... in binary, which has 1s at even positions
	return n&0x55555555 != 0
}

func main() {
	// Test cases
	testCases := []int{16, 5, 1, 0, -1, 8, 64, 256}

	fmt.Println("Testing isPowerOfFour:")
	for _, n := range testCases {
		result := isPowerOfFour(n)
		fmt.Printf("isPowerOfFour(%d) = %t\n", n, result)
	}

	fmt.Println("\nTesting isPowerOfFourBitwise:")
	for _, n := range testCases {
		result := isPowerOfFourBitwise(n)
		fmt.Printf("isPowerOfFourBitwise(%d) = %t\n", n, result)
	}
}
