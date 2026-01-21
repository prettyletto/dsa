package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic case",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "unordered result acceptable",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "negative numbers",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
		{
			name:   "no solution",
			nums:   []int{1, 2, 3},
			target: 7,
			want:   []int{},
		},
		{
			name:   "same number used once",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v",
					tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
