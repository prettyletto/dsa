package arrays

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "basic case",
			nums:   []int{3},
			target: 3,
			want:   1,
		},
		{
			name:   "single element",
			nums:   []int{3},
			target: 2,
			want:   0,
		},
		{
			name:   "negative number",
			nums:   []int{1, -1, 0},
			target: 0,
			want:   3,
		},
		{
			name:   "another test",
			nums:   []int{2, 3, 1},
			target: 6,
			want:   1,
		},
	}

	for _, t := range tests {
		if subarraySum(t.nums, t.target) != t.want {
			panic("test failed")
		}
	}
}
