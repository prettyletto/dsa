package arrays

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "basic example",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5, // buy at 1, sell at 6
		},
		{
			name:   "monotonically decreasing",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "single element",
			prices: []int{5},
			want:   0,
		},
		{
			name:   "two elements profit",
			prices: []int{1, 5},
			want:   4,
		},
		{
			name:   "two elements loss",
			prices: []int{5, 1},
			want:   0,
		},
		{
			name:   "multiple dips",
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   4, // buy at 2, sell at 6
		},
		{
			name:   "flat prices",
			prices: []int{2, 2, 2, 2},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
