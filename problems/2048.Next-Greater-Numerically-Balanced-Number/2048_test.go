package problem2048

import "testing"

func TestNextBeautifulNumber(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 22},
		{1000, 1333},
		{3000, 3133},
		{0, 1},
		{10, 22},
		{100, 122},
		{1234, 1333},
		{5000, 14444},
		{10000, 14444},
		{100000, 122333},
		{999999, 1224444},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := nextBeautifulNumber(tt.n)
			if got != tt.want {
				t.Errorf("nextBeautifulNumber(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
