package leetcode

import (
	"testing"
)

func TestSortVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "lEetcOde",
			want: "lEOtcede",
		},
		{
			name: "Example 2",
			s:    "lYmpH",
			want: "lYmpH",
		},
		{
			name: "Empty string",
			s:    "",
			want: "",
		},
		{
			name: "Only vowels",
			s:    "aEiOu",
			want: "EOaiu",
		},
		{
			name: "Only consonants",
			s:    "bcdfg",
			want: "bcdfg",
		},
		{
			name: "Mixed case vowels",
			s:    "AeIoU",
			want: "AIUeo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortVowels(tt.s)
			if got != tt.want {
				t.Errorf("sortVowels(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
