package main

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Power of 4: 1", 1, true},                         // 4^0 = 1
		{"Power of 4: 4", 4, true},                         // 4^1 = 4
		{"Power of 4: 16", 16, true},                       // 4^2 = 16
		{"Power of 4: 64", 64, true},                       // 4^3 = 64
		{"Power of 4: 256", 256, true},                     // 4^4 = 256
		{"Power of 4: 1024", 1024, true},                   // 4^5 = 1024
		{"Not power of 4: 2", 2, false},                    // 2^1
		{"Not power of 4: 8", 8, false},                    // 2^3
		{"Not power of 4: 5", 5, false},                    // Not a power of 2
		{"Not power of 4: 6", 6, false},                    // Not a power of 2
		{"Not power of 4: 7", 7, false},                    // Not a power of 2
		{"Not power of 4: 9", 9, false},                    // Not a power of 2
		{"Not power of 4: 10", 10, false},                  // Not a power of 2
		{"Not power of 4: 12", 12, false},                  // Not a power of 2
		{"Not power of 4: 15", 15, false},                  // Not a power of 2
		{"Not power of 4: 17", 17, false},                  // Not a power of 2
		{"Zero: 0", 0, false},                              // 0 is not positive
		{"Negative: -1", -1, false},                        // Negative numbers are not powers of 4
		{"Negative: -4", -4, false},                        // Negative numbers are not powers of 4
		{"Large power of 4: 1073741824", 1073741824, true}, // 4^15
		{"Large non-power: 1073741825", 1073741825, false}, // 4^15 + 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPowerOfFour(tt.input)
			if result != tt.expected {
				t.Errorf("isPowerOfFour(%d) = %t, want %t", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPowerOfFourBitwise(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Power of 4: 1", 1, true},                         // 4^0 = 1
		{"Power of 4: 4", 4, true},                         // 4^1 = 4
		{"Power of 4: 16", 16, true},                       // 4^2 = 16
		{"Power of 4: 64", 64, true},                       // 4^3 = 64
		{"Power of 4: 256", 256, true},                     // 4^4 = 256
		{"Power of 4: 1024", 1024, true},                   // 4^5 = 1024
		{"Not power of 4: 2", 2, false},                    // 2^1
		{"Not power of 4: 8", 8, false},                    // 2^3
		{"Not power of 4: 5", 5, false},                    // Not a power of 2
		{"Not power of 4: 6", 6, false},                    // Not a power of 2
		{"Not power of 4: 7", 7, false},                    // Not a power of 2
		{"Not power of 4: 9", 9, false},                    // Not a power of 2
		{"Not power of 4: 10", 10, false},                  // Not a power of 2
		{"Not power of 4: 12", 12, false},                  // Not a power of 2
		{"Not power of 4: 15", 15, false},                  // Not a power of 2
		{"Not power of 4: 17", 17, false},                  // Not a power of 2
		{"Zero: 0", 0, false},                              // 0 is not positive
		{"Negative: -1", -1, false},                        // Negative numbers are not powers of 4
		{"Negative: -4", -4, false},                        // Negative numbers are not powers of 4
		{"Large power of 4: 1073741824", 1073741824, true}, // 4^15
		{"Large non-power: 1073741825", 1073741825, false}, // 4^15 + 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPowerOfFourBitwise(tt.input)
			if result != tt.expected {
				t.Errorf("isPowerOfFourBitwise(%d) = %t, want %t", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests to compare performance
func BenchmarkIsPowerOfFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPowerOfFour(1073741824) // Large power of 4
	}
}

func BenchmarkIsPowerOfFourBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPowerOfFourBitwise(1073741824) // Large power of 4
	}
}

// Test that both functions give the same results
func TestConsistency(t *testing.T) {
	testNumbers := []int{-100, -16, -4, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 16, 17, 31, 32, 33, 63, 64, 65, 255, 256, 257, 1023, 1024, 1025}

	for _, n := range testNumbers {
		result1 := isPowerOfFour(n)
		result2 := isPowerOfFourBitwise(n)
		if result1 != result2 {
			t.Errorf("Inconsistent results for %d: isPowerOfFour=%t, isPowerOfFourBitwise=%t", n, result1, result2)
		}
	}
}
